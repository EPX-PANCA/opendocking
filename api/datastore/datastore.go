package datastore

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	portainer "github.com/opendocking/opendocking/api"
	"github.com/opendocking/opendocking/api/dataservices"
	portainerErrors "github.com/opendocking/opendocking/api/dataservices/errors"

	"github.com/rs/zerolog/log"
)

// NewStore initializes a new Store and the associated services
func NewStore(cliFlags *portainer.CLIFlags, fileService portainer.FileService, connection portainer.Connection) *Store {
	return &Store{
		flags:       cliFlags,
		fileService: fileService,
		connection:  connection,
	}
}

// Open opens and initializes the BoltDB database.
func (store *Store) Open() (newStore bool, err error) {
	encryptionReq, err := store.connection.NeedsEncryptionMigration()
	if err != nil {
		return false, err
	}

	if encryptionReq {
		// NeedsEncryptionMigration() sets encrypted=true as a side effect when a key exists.
		// We need to set it back to false so GetDatabaseFilePath() returns the path to the
		// actual unencrypted file (portainer.db) that we want to back up.
		if err := store.connection.SetEncrypted(false); err != nil {
			return false, err
		}

		// Use backupDBFile directly since connection isn't open yet
		// and we don't want to trigger the close/open cycle of Backup()
		backupFilename, err := store.backupDBFile("")
		if err != nil {
			return false, fmt.Errorf("failed to backup database prior to encrypting: %w", err)
		}

		if err := store.encryptDB(); err != nil {
			innerErr := store.RestoreFromFile(backupFilename) // restore from backup if encryption fails
			return false, errors.Join(err, innerErr)
		}
	}

	if err := store.connection.Open(); err != nil {
		return false, err
	}

	if err := store.initServices(); err != nil {
		return false, err
	}

	// If no settings object exists then assume we have a new store
	if _, err := store.SettingsService.Settings(); err != nil {
		if store.IsErrObjectNotFound(err) {
			return true, nil
		}

		return false, err
	}

	return false, nil
}

func (store *Store) Close() error {
	return store.connection.Close()
}

func (store *Store) UpdateTx(fn func(dataservices.DataStoreTx) error) error {
	return store.connection.UpdateTx(func(tx portainer.Transaction) error {
		return fn(&StoreTx{store: store, tx: tx})
	})
}

func (store *Store) ViewTx(fn func(dataservices.DataStoreTx) error) error {
	return store.connection.ViewTx(func(tx portainer.Transaction) error {
		return fn(&StoreTx{store: store, tx: tx})
	})
}

// BackupTo backs up db to a provided writer.
// It does hot backup and doesn't block other database reads and writes
func (store *Store) BackupTo(w io.Writer) error {
	return store.connection.BackupTo(w)
}

// CheckCurrentEdition checks if the edition is compatible with OpenDocking
// OpenDocking accepts databases from OpenDocking, BE, and EE editions
func (store *Store) CheckCurrentEdition() error {
	edition := store.edition()
	if edition != portainer.OpenDocking &&
		edition != portainer.PortainerCE &&
		edition != portainer.PortainerBE &&
		edition != portainer.PortainerEE {
		return portainerErrors.ErrWrongDBEdition
	}

	return nil
}

func (store *Store) edition() portainer.SoftwareEdition {
	edition, err := store.VersionService.Edition()
	if store.IsErrObjectNotFound(err) {
		edition = portainer.OpenDocking
	}

	return edition
}

// TODO: move the use of this to dataservices.IsErrObjectNotFound()?
func (store *Store) IsErrObjectNotFound(e error) bool {
	return errors.Is(e, portainerErrors.ErrObjectNotFound)
}

func (store *Store) Connection() portainer.Connection {
	return store.connection
}

func (store *Store) Rollback(force bool) error {
	return store.connectionRollback(force)
}

func (store *Store) encryptDB() error {
	if err := store.connection.SetEncrypted(false); err != nil {
		return err
	}

	if err := store.connection.Open(); err != nil {
		return err
	}

	if err := store.initServices(); err != nil {
		return err
	}

	// The DB is not currently encrypted.  First save the encrypted db filename
	oldFilename := store.connection.GetDatabaseFilePath()
	log.Info().Msg("encrypting database")

	// export file path for backup
	exportFilename := path.Join(store.databasePath() + "." + fmt.Sprintf("backup-%d.json", time.Now().Unix()))

	log.Info().Str("filename", exportFilename).Msg("exporting database backup")

	if err := store.Export(exportFilename); err != nil {
		log.Error().Str("filename", exportFilename).Err(err).Msg("failed to export")

		return err
	}

	log.Info().Msg("database backup exported")

	// Close existing un-encrypted db so that we can delete the file later
	if err := store.connection.Close(); err != nil {
		return err
	}

	if err := store.Import(exportFilename); err != nil {
		log.Error().Err(err).Msg("failed to import database backup")

		// Remove the new encrypted file that we failed to import
		if err := os.Remove(store.connection.GetDatabaseFilePath()); err != nil {
			log.Error().Msg("failed to remove the file after import failure")
		}

		log.Fatal().Err(portainerErrors.ErrDBImportFailed).Msg("")
	}

	if err := os.Remove(oldFilename); err != nil {
		log.Error().Msg("failed to remove the un-encrypted db file")
	}

	if err := os.Remove(exportFilename); err != nil {
		log.Error().Msg("failed to remove the json backup file")
	}

	// Close db connection
	if err := store.connection.Close(); err != nil {
		return err
	}

	log.Info().Msg("database successfully encrypted")

	return nil
}
