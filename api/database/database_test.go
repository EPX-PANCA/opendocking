package database

import (
	"testing"

	"github.com/opendocking/opendocking/api/database/boltdb"
	"github.com/opendocking/opendocking/api/filesystem"

	"github.com/stretchr/testify/require"
)

func TestNewDatabase(t *testing.T) {
	t.Parallel()
	dbPath := filesystem.JoinPaths(t.TempDir(), "test.db")
	connection, err := NewDatabase("boltdb", dbPath, nil, false)
	require.NoError(t, err)
	require.NotNil(t, connection)

	_, ok := connection.(*boltdb.DbConnection)
	require.True(t, ok)

	connection, err = NewDatabase("unknown", dbPath, nil, false)
	require.Error(t, err)
	require.Nil(t, connection)
}
