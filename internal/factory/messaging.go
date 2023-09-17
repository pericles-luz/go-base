package factory

import (
	"github.com/pericles-luz/go-base/internal/migration"
	"github.com/pericles-luz/go-base/pkg/infra/database"
)

func NewMessageDB(pool *database.Pool) *migration.MessageDB {
	return migration.NewMessageDB(pool.GetConnection(migration.DATABASE_MESSAGING).GetDatabase())
}

func NewMessageService(pool *database.Pool) *migration.MessageService {
	persistence := NewMessageDB(pool)
	return migration.NewMessageService(persistence)
}
