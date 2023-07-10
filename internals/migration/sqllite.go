package migration

import "database/sql"

func CreateTableSQLite(db *sql.DB) error {
	if _, err := db.Exec(`
	create table if not exists RabbitCache(RabbitCacheID string primary key, DE_Exchange string, DE_RoutingKey string, JS_Data text, SN_Durable integer, TS_Operacao string, ID_Status integer default 0)
	`); err != nil {
		return err
	}
	return nil
}
