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
	conn := pool.GetConnection(db.GetName())
	require.Equal(t, db, conn)
	conn = pool.GetConnection("test2")
	require.Equal(t, db, conn)
}

func TestPoolGetConnectionMustFailIfThereIsNoConnection(t *testing.T) {
	pool := database.NewPool()
	require.Nil(t, pool.GetConnection("test"))
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
	require.Nil(t, pool.GetConnection(db.GetName()))
}
