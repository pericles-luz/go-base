package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Database struct {
	file     ConfigBase
	DBName   string `json:"dbname,omitempty"`
	Engine   string `json:"engine,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (cfg *Database) Validate() error {
	if cfg.Engine == "sqlite3" {
		return nil
	}
	if cfg.Username == "" {
		return errors.New("database username not informed")
	}
	if cfg.Password == "" {
		return errors.New("database username password not informed")
	}
	return nil
}

func (cfg *Database) GetDSN() (string, error) {
	if err := cfg.Validate(); err != nil {
		return "", cfg.Validate()
	}
	if cfg.Engine == "sqlite3" {
		return ":memory:", nil
	}
	if cfg.Engine == "" {
		log.Println("Engine assumed mysql")
		cfg.Engine = "mysql"
	}
	if cfg.Host == "" {
		log.Println("Host assumed mariadb")
		cfg.Host = "127.0.0.1"
	}
	if cfg.DBName == "" {
		log.Println("DBName assumed buscaCompleta")
		cfg.DBName = "mariadb"
	}
	if cfg.Port == 0 {
		log.Println("Port assumed 3306")
		cfg.Port = 3306
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	), nil
}

func NewDatabase(file string) (*Database, error) {
	result := &Database{}
	jsonConfiguration, err := result.file.ReadConfigurationFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonConfiguration, &result)
	if err != nil {
		return nil, err
	}
	if result.Validate() != nil {
		return nil, result.Validate()
	}
	return result, nil
}
