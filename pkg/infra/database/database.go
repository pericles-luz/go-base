package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pericles-luz/go-base/pkg/conf"
)

const (
	ENGINE_SQLITE   = "sqlite3"
	ENGINE_MYSQL    = "mysql"
	ENGINE_POSTGRES = "postgres"
)

type IDatabase interface {
	connect(*conf.Database) error
	Exec(sql string, data ...interface{}) error
	GetOne(sql string, data ...interface{}) ([]byte, error)
	GetRecord(sql string, data ...interface{}) (map[string]interface{}, error)
	GetRecords(sql string, data ...interface{}) ([]map[string]interface{}, error)
	Insert(tableName string, data map[string]interface{}) error
	Update(tableName string, data map[string]interface{}) error
	IsConnected() bool
	IsSQLite() bool
	IsMySQL() bool
	IsPostgres() bool
	GetLastInsertId() uint64
	GetResult() (sql.Result, error)
	GetDatabase() *sql.DB
	GetName() string
	SetEngine(engine string)
}

type Database struct {
	database *sql.DB
	res      sql.Result
	engine   string
	name     string
}

func NewDatabase(cfg *conf.Database) (*Database, error) {
	database := Database{}
	if cfg == nil {
		return nil, errors.New("no configuration provided")
	}
	if err := database.connect(cfg); nil != err {
		log.Println("database not initialized")
		return nil, err
	}
	if cfg.Name == "" {
		cfg.Name = cfg.DBName
	}
	database.name = cfg.Name
	database.SetEngine(ENGINE_MYSQL)
	if cfg.Engine != "" {
		database.SetEngine(cfg.Engine)
	}
	return &database, nil
}

func NewDatabaseWithConnection(db *sql.DB) Database {
	database := Database{}
	database.database = db
	return database
}

func (db *Database) connect(cfg *conf.Database) error {
	db.engine = cfg.Engine
	log.Println("database engine: " + db.engine)
	err := db.openDBWithStartupWait(cfg)
	if err != nil {
		log.Fatal("database not available")
		return errors.New("database not available")
	}
	db.configureConnectionPool()
	return err
}

func (db *Database) GetDatabase() *sql.DB {
	return db.database
}

func (db *Database) openDBWithStartupWait(cfg *conf.Database) error {
	var startupTimeout = func() time.Duration {
		str := "10s"
		d, err := time.ParseDuration(str)
		if err != nil {
			log.Println("db startup timed out")
		}
		return d
	}()
	startupDeadline := time.Now().Add(startupTimeout)
	for {
		if time.Now().After(startupDeadline) {
			log.Fatal("database did not start up within")
			return errors.New("database did not start up within")
		}
		err := db.open(cfg)
		if err == nil {
			err = db.database.Ping()
		}
		if err != nil {
			log.Println("ping failed: ", err)
			time.Sleep(startupTimeout / 10)
			continue
		}
		return err
	}
}

// Open creates a new DB handle with the given schema by connecting to
// the database identified by dataSource (e.g., "dbname=mypgdb" or
// blank to use the PG* env vars).
//
// Open assumes that the database already exists.
func (db *Database) open(Database *conf.Database) error {
	dsn, err := Database.GetDSN()
	if nil != err {
		log.Fatal("dsn not generated")
		return errors.New("dsn not generated")
	}
	db.database, err = sql.Open(Database.Engine, dsn)
	if err != nil {
		log.Println(err.Error())
		log.Fatal("database not openned")
		return errors.New("database not openned")
	}
	return err
}

// Ping attempts to contact the database and returns a non-nil error upon failure. It is intended to
// be used by health checks.
//func Ping(ctx context.Context) error { return Global.PingContext(ctx) }

// configureConnectionPool sets reasonable sizes on the built in DB queue. By
// default the connection pool is unbounded, which leads to the error `pq:
// sorry too many clients already`.
func (db *Database) configureConnectionPool() {
	maxOpen := 30
	if db.IsSQLite() {
		maxOpen = 1
	}
	db.database.SetMaxOpenConns(maxOpen)
	db.database.SetMaxIdleConns(maxOpen)
	db.database.SetConnMaxLifetime(time.Minute)
}

func (db *Database) IsConnected() bool {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err := db.database.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return false
	}
	return nil == err
}

func (db *Database) Close() error {
	err := db.database.Close()
	if nil != err {
		log.Print("failed to close database", err.Error())
	}
	return err
}

func (db *Database) GetResult() (sql.Result, error) {
	if nil == db.res {
		return nil, errors.New("no result to return")
	}
	return db.res, nil
}

func (db *Database) GetLastInsertId() uint64 {
	if db, err := db.GetResult(); err == nil {
		if id, err := db.LastInsertId(); err == nil {
			return uint64(id)
		}
	}
	return 0
}

func (db *Database) GetStmt(sql string) (*sql.Stmt, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	stmt, err := db.database.PrepareContext(ctx, sql)
	if nil != err {
		cancelfunc()
		log.Println("failed preparing sql request", err.Error(), sql)
		return nil, err
	}
	cancelfunc()
	return stmt, err
}

func (db *Database) Exec(sql string, data ...interface{}) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.GetStmt(sql)
	if nil != err {
		log.Println("failed preparing sql request", err.Error(), sql, data)
		return err
	}
	defer stmt.Close()
	if len(data) == 0 {
		_, err := stmt.ExecContext(ctx)
		if err != nil {
			log.Printf("Error %s executing statement\n", err)
			return err
		}
	} else {
		res, err := stmt.ExecContext(ctx, data...)
		if err != nil {
			log.Printf("Error %s executing statement\n", err)
			return err
		}
		db.res = res
	}
	return nil
}

func (db *Database) GetOne(sql string, data ...interface{}) ([]byte, error) {
	if !strings.HasSuffix(sql, "LIMIT 1") {
		sql = fmt.Sprintf("%s LIMIT 1", sql)
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.GetStmt(sql)
	if nil != err {
		log.Println("failed preparing sql request", err.Error(), sql)
		return nil, err
	}
	defer stmt.Close()

	query, err := stmt.QueryContext(ctx, data...)
	if err != nil {
		log.Printf("Error %s executing statement\n", err)
		return nil, err
	}
	if !query.Next() {
		return nil, nil
	}
	var result []byte
	if err = query.Scan(&result); err != nil {
		log.Println("failed loading data", err.Error())
		return nil, err
	}
	return result, nil
}

func (db *Database) makeRecord(query *sql.Rows) (map[string]interface{}, error) {
	var columns []string
	var err error
	if columns, err = query.Columns(); err != nil {
		log.Println("failed reading columns", err.Error())
		return nil, err
	}
	line := make([]sql.NullString, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range line {
		scanArgs[i] = &line[i]
	}
	if err := query.Scan(scanArgs...); err != nil {
		log.Println("fail scanning record", err.Error())
		return nil, err
	}
	result := make(map[string]interface{})
	for i := range line {
		if line[i].Valid {
			result[columns[i]] = line[i].String
			continue
		}
		result[columns[i]] = nil
	}
	return result, nil
}

func (db *Database) GetRecord(sqlString string, data ...interface{}) (map[string]interface{}, error) {
	if !strings.HasSuffix(sqlString, "LIMIT 1") {
		sqlString = fmt.Sprintf("%s LIMIT 1", sqlString)
	}
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.GetStmt(sqlString)
	if nil != err {
		log.Println("failed preparing sql request", err.Error(), sqlString)
		return nil, err
	}
	defer stmt.Close()
	query, err := stmt.QueryContext(ctx, data...)
	if err != nil {
		log.Printf("error %s executing statement\n", err)
		return nil, err
	}
	if err := query.Err(); err == sql.ErrNoRows {
		log.Println("no record found", err.Error())
		return nil, err
	}
	if err := query.Err(); err != nil {
		log.Println("error reading next record", err.Error())
		return nil, err
	}
	if query == nil {
		log.Println("query is nil")
		return nil, errors.New("query is nil")
	}
	if !query.Next() {
		return nil, nil
	}
	return db.makeRecord(query)
}

func (db *Database) GetRecords(sqlString string, data ...interface{}) ([]map[string]interface{}, error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelfunc()
	stmt, err := db.GetStmt(sqlString)
	if nil != err {
		log.Println("failed preparing sql request", err.Error(), sqlString)
		return nil, err
	}
	defer stmt.Close()
	query, err := stmt.QueryContext(ctx, data...)
	if err != nil {
		log.Printf("error %s executing statement\n", err)
		return nil, err
	}
	if err := query.Err(); err == sql.ErrNoRows {
		log.Println("no records found", err.Error())
		return nil, err
	}
	if err := query.Err(); err != nil {
		log.Println("error reading next record", err.Error())
		return nil, err
	}
	if query == nil {
		log.Println("query is nil")
		return nil, errors.New("query is nil")
	}
	result := make([]map[string]interface{}, 0)
	var line map[string]interface{}
	for query.Next() {
		if line, err = db.makeRecord(query); nil != err {
			log.Println("error reading data", err.Error())
			return nil, err
		}
		result = append(result, line)
	}
	return result, nil
}

func (db *Database) Insert(tableName string, data map[string]interface{}) error {
	var keys []string
	var values []interface{}
	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(keys, ","), strings.Repeat("?,", len(keys)-1)+"?")
	err := db.Exec(sql, values...)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Update(tableName string, data map[string]interface{}) error {
	keyName := tableName + "UD"
	if data[keyName] == nil && data[tableName+"ID"] != nil {
		keyName = tableName + "ID"
	}
	if data[keyName] == nil {
		return errors.New("no id provided for update")
	}
	var keys []string
	var values []interface{}
	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s=?", tableName, strings.Join(keys, "=?,")+"=?", keyName)
	values = append(values, data[keyName])
	err := db.Exec(sql, values...)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) SetEngine(engine string) {
	db.engine = engine
}

func (db *Database) IsMySQL() bool {
	return db.engine == ENGINE_MYSQL
}

func (db *Database) IsPostgres() bool {
	return db.engine == ENGINE_POSTGRES
}

func (db *Database) IsSQLite() bool {
	return db.engine == ENGINE_SQLITE
}

func (db *Database) GetName() string {
	return db.name
}
