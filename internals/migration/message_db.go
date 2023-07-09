package migration

import (
	"database/sql"
)

type MessageDB struct {
	db *sql.DB
}

func NewMessageDB(db *sql.DB) *MessageDB {
	return &MessageDB{db: db}
}

func (p *MessageDB) Get(id string) (*Message, error) {
	var message Message
	stmt, err := p.db.Prepare(`select * from RabbitCache where RabbitCacheID=?`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(
		&message.RabbitCacheID,
		&message.DE_Exchange,
		&message.DE_RoutingKey,
		&message.JS_Data,
		&message.SN_Durable,
		&message.TS_Operacao,
	)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (p *MessageDB) GetNext() (*Message, error) {
	var message Message
	var createdAt sql.NullString
	stmt, err := p.db.Prepare(`select RabbitCacheID, DE_Exchange, DE_RoutingKey, JS_Data, SN_Durable, TS_Operacao from RabbitCache order by RabbitCacheID limit 1`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow().Scan(
		&message.RabbitCacheID,
		&message.DE_Exchange,
		&message.DE_RoutingKey,
		&message.JS_Data,
		&message.SN_Durable,
		&createdAt,
	)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (p *MessageDB) Save(message *Message) error {
	stmt, err := p.db.Prepare(`insert into RabbitCache(RabbitCacheID, DE_Exchange, DE_RoutingKey, JS_Data, SN_Durable) values(?,?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		message.RabbitCacheID,
		message.DE_Exchange,
		message.DE_RoutingKey,
		message.JS_Data,
		message.SN_Durable,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *MessageDB) Delete(id string) error {
	stmt, err := p.db.Prepare(`delete from RabbitCache where RabbitCacheID=?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
