package factory

import (
	"github.com/pericles-luz/go-base/internals/migration"
	"github.com/pericles-luz/go-base/pkg/infra/database"
)

func MewMessageDB(pool *database.Pool) *migration.MessageDB {
	return migration.NewMessageDB(pool.GetConnection(migration.DATABASE_MESSAGING).GetDatabase())
}
