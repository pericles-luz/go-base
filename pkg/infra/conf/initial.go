package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Config for the environment, read from a json file
type Config struct {
	file                    ConfigBase
	Debug                   bool
	Production              bool
	DBConfigurationFile     string
	AgnuDBConfigurationFile string
	DBConfiguration         *DBConfiguration
	AgnuDBConfiguration     *DBConfiguration
	JwtSecret               string
}

func (cfg *Config) Validate() error {
	if cfg.DBConfigurationFile == "" {
		log.Println("empty file name:", cfg.DBConfigurationFile)
		return errors.New("no db configuration defined")
	}
	return nil
}

type DBConfiguration struct {
	DBName   string `json:"dbname,omitempty"`
	Engine   string `json:"engine,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (cfg *DBConfiguration) Validate() error {
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

func (cfg *DBConfiguration) GetDSN() (string, error) {
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

func (cfg *Config) parseDBConfiguration() error {
	if err := cfg.Validate(); err != nil {
		return err
	}
	if cfg.DBConfiguration != nil {
		log.Println("DBConfiguration already loaded")
		return nil
	}
	dBConfiguration := &DBConfiguration{}
	jsonConfiguration, err := cfg.file.ReadConfigurationFile(cfg.DBConfigurationFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonConfiguration, &dBConfiguration)
	if err != nil {
		return err
	}
	if dBConfiguration.Validate() != nil {
		return dBConfiguration.Validate()
	}
	cfg.DBConfiguration = dBConfiguration
	return nil
}

func (cfg *Config) parseAgnuDBConfiguration() error {
	if err := cfg.Validate(); err != nil {
		log.Println("parseAgnuDBConfiguration: ", err)
		return err
	}
	if cfg.AgnuDBConfigurationFile == "" {
		log.Println("empty file name:", cfg.AgnuDBConfigurationFile)
		return nil
	}
	if cfg.AgnuDBConfiguration != nil {
		log.Println("AgnuDBConfiguration already loaded")
		return nil
	}
	dBConfiguration := &DBConfiguration{}
	jsonConfiguration, err := cfg.file.ReadConfigurationFile(cfg.AgnuDBConfigurationFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonConfiguration, &dBConfiguration)
	if err != nil {
		return err
	}
	if dBConfiguration.Validate() != nil {
		return dBConfiguration.Validate()
	}
	cfg.AgnuDBConfiguration = dBConfiguration
	return nil
}

// NewInitialConfig reads configuration from json file and validates it
func NewInitialConfig(fileName string) (Config, error) {
	cfg := Config{}
	jsonConfiguration, err := cfg.file.ReadConfigurationFile(fileName)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(jsonConfiguration, &cfg)
	if err != nil {
		return cfg, err
	}
	if err = cfg.Validate(); err != nil {
		return cfg, err
	}
	if err = cfg.parseDBConfiguration(); err != nil {
		return cfg, err
	}
	if err = cfg.parseAgnuDBConfiguration(); err != nil {
		return cfg, err
	}
	return cfg, nil
}
