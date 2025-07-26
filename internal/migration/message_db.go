package migration

import (
	"context"
	"database/sql"
	"time"
)

const (
	STATUS_PENDING = 0
	STATUS_SENDING = 1
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
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(
		&message.RabbitCacheID,
		&message.DE_Exchange,
		&message.DE_RoutingKey,
		&message.JS_Data,
		&message.SN_Durable,
		&message.TS_Operacao,
	)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (p *MessageDB) GetNext() (*Message, error) {
	var message Message
	var createdAt, id, exchange, routing, data sql.NullString
	var durable sql.NullInt64
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	stmt, err := p.db.PrepareContext(ctx, `select RabbitCacheID, DE_Exchange, DE_RoutingKey, JS_Data, SN_Durable, TS_Operacao from RabbitCache order by ID_Status, RabbitCacheID limit 1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(ctx).Scan(
		&id,
		&exchange,
		&routing,
		&data,
		&durable,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}
	if !id.Valid {
		return nil, sql.ErrNoRows
	}
	if !exchange.Valid || !routing.Valid || !data.Valid {
		return nil, sql.ErrNoRows
	}
	if !durable.Valid {
		durable.Int64 = 1 // Default to 1 if not set
	}
	if !createdAt.Valid {
		createdAt.String = time.Now().Format("2006-01-02 15:04:05")
	}
	message.TS_Operacao, _ = time.Parse("2006-01-02 15:04:05", createdAt.String)
	message.RabbitCacheID = id.String
	message.DE_Exchange = exchange.String
	message.DE_RoutingKey = routing.String
	message.JS_Data = data.String
	message.SN_Durable = int16(durable.Int64)
	if message.SN_Durable == 0 {
		message.SN_Durable = 1 // Default to 1 if not set
	}
	err = p.SetStatus(message.RabbitCacheID, STATUS_SENDING)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (p *MessageDB) Save(message *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	stmt, err := p.db.PrepareContext(ctx, `insert into RabbitCache(RabbitCacheID, DE_Exchange, DE_RoutingKey, JS_Data, SN_Durable) values(?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx,
		message.RabbitCacheID,
		message.DE_Exchange,
		message.DE_RoutingKey,
		message.JS_Data,
		message.SN_Durable,
	)
	return err
}

func (m *MessageDB) SetStatus(id string, status int) error {
	stmt, err := m.db.Prepare(`update RabbitCache set ID_Status=? where RabbitCacheID=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(status, id)
	return err
}

func (p *MessageDB) Delete(id string) error {
	stmt, err := p.db.Prepare(`delete from RabbitCache where RabbitCacheID=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
