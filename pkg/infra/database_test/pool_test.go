package database_test

import (
	"testing"

	"github.com/pericles-luz/go-base/pkg/infra/database"
	"github.com/stretchr/testify/require"
)

func TestPoolAddConnectionsCorrectly(t *testing.T) {
	pool := database.NewPool()
	db, err := databaseConnection(t)
	require.NoError(t, err)
	defer db.Close()
	require.True(t, db.IsConnected())
	require.NoError(t, pool.AddConnection(db.GetName(), db))
	require.NoError(t, pool.AddConnection("test2", db))
	conn, err := pool.GetConnection(db.GetName())
	require.NoError(t, err)
	require.Equal(t, db, conn)
	conn, err = pool.GetConnection("test2")
	require.NoError(t, err)
	require.Equal(t, db, conn)
}

func TestPoolGetConnectionMustFailIfThereIsNoConnection(t *testing.T) {
	pool := database.NewPool()
	_, err := pool.GetConnection("test")
	require.Error(t, err)
}

func TestPoolDeleteConnectionMustFailIfThereIsNoConnection(t *testing.T) {
	pool := database.NewPool()
	err := pool.RemoveConnection("test")
	require.Error(t, err)
}

func TestPoolGetConnectionMustFailAfterDeleteConnection(t *testing.T) {
	pool := database.NewPool()
	db, err := databaseConnection(t)
	require.NoError(t, err)
	defer db.Close()
	require.NoError(t, pool.AddConnection(db.GetName(), db))
	require.NoError(t, pool.RemoveConnection(db.GetName()))
	_, err = pool.GetConnection(db.GetName())
	require.Error(t, err)
}
