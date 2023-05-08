package migration

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pericles-luz/go-base/pkg/infra/database"
	"github.com/stretchr/testify/require"
)

const (
	DATABASE_GENERAL   = "general"
	DATABASE_FINANCE   = "finance"
	DATABASE_MESSAGING = "messaging"
)

func SetupTest(t *testing.T) (func(t *testing.T), *database.Pool) {
	db := GetDB()
	require.NotNil(t, db)
	// evita erro de "table does not exist"
	db.SetMaxOpenConns(1)
	require.NotNil(t, db)
	require.NoError(t, db.Ping())
	require.NoError(t, CreateTableSQLite(db))
	require.NoError(t, GenerateTestData(db))
	pool := GetPool(db)
	return func(t *testing.T) {
		// t.Log("teardown test")
	}, pool
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	return db
}

func GetPool(db *sql.DB) *database.Pool {
	pool := database.NewPool()
	data := database.NewDatabaseWithConnection(db)
	data.SetEngine(database.ENGINE_SQLITE)
	pool.AddConnection(DATABASE_GENERAL, &data)
	pool.AddConnection(DATABASE_FINANCE, &data)
	pool.AddConnection(DATABASE_MESSAGING, &data)
	return pool
}
