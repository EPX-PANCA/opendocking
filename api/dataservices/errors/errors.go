package errors

import (
	"errors"
)

var (
	ErrObjectNotFound     = errors.New("object not found inside the database")
	ErrWrongDBEdition     = errors.New("the database edition is not compatible with this version of OpenDocking")
	ErrDBImportFailed     = errors.New("importing backup failed")
	ErrDatabaseIsUpdating = errors.New("database is currently in updating state. Failed prior upgrade. Please restore from backup or delete the database and restart OpenDocking")
)
