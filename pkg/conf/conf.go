package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
)

// Config for the environment, read from a json file
type Config struct {
	file                      ConfigBase
	Debug                     bool
	Production                bool
	DbConfigurationFile       string
	ZimpeDbConfigurationFile  string
	VerzDbConfigurationFile   string
	DialerDbConfigurationFile string
	DbConfiguration           *DbConfiguration
	ZimpeDbConfiguration      *DbConfiguration
	VerzDbConfiguration       *DbConfiguration
	DialerDbConfiguration     *DbConfiguration
	JwtSecret                 string
}

func (cfg *Config) Validate() error {
	if cfg.DbConfigurationFile == "" {
		log.Println("empty file name:", cfg.DbConfigurationFile)
		return errors.New("no db configuration defined")
	}
	return nil
}

type DbConfiguration struct {
	DBName   string `json:"dbname,omitempty"`
	Engine   string `json:"engine,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (cfg *DbConfiguration) Validate() error {
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

func (cfg *DbConfiguration) GetDSN() (string, error) {
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
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	), nil
}

func (cfg *Config) parseDbConfiguration() error {
	if err := cfg.Validate(); err != nil {
		return err
	}
	if cfg.DbConfiguration != nil {
		log.Println("DbConfiguration already loaded")
		return nil
	}
	dbConfiguration := &DbConfiguration{}
	jsonConfiguration, err := cfg.file.ReadConfigurationFile(cfg.DbConfigurationFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonConfiguration, &dbConfiguration)
	if err != nil {
		return err
	}
	if dbConfiguration.Validate() != nil {
		return dbConfiguration.Validate()
	}
	cfg.DbConfiguration = dbConfiguration
	return nil
}

func (cfg *Config) parseZimpeDbConfiguration() error {
	if err := cfg.Validate(); err != nil {
		log.Println("parseZimpeDbConfiguration: ", err)
		return err
	}
	if cfg.ZimpeDbConfigurationFile == "" {
		log.Println("empty file name:", cfg.ZimpeDbConfigurationFile)
		return nil
	}
	if cfg.ZimpeDbConfiguration != nil {
		log.Println("ZimpeDbConfiguration already loaded")
		return nil
	}
	dbConfiguration := &DbConfiguration{}
	jsonConfiguration, err := cfg.file.ReadConfigurationFile(cfg.ZimpeDbConfigurationFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonConfiguration, &dbConfiguration)
	if err != nil {
		return err
	}
	if dbConfiguration.Validate() != nil {
		return dbConfiguration.Validate()
	}
	cfg.ZimpeDbConfiguration = dbConfiguration
	return nil
}

func (cfg *Config) parseVerzDbConfiguration() error {
	if err := cfg.Validate(); err != nil {
		log.Println("parseVerzDbConfiguration: ", err)
		return err
	}
	if cfg.VerzDbConfigurationFile == "" {
		log.Println("empty file name:", cfg.VerzDbConfigurationFile)
		return nil
	}
	if cfg.VerzDbConfiguration != nil {
		log.Println("VerzDbConfiguration already loaded")
		return nil
	}
	dbConfiguration := &DbConfiguration{}
	jsonConfiguration, err := cfg.file.ReadConfigurationFile(cfg.VerzDbConfigurationFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonConfiguration, &dbConfiguration)
	if err != nil {
		return err
	}
	if dbConfiguration.Validate() != nil {
		return dbConfiguration.Validate()
	}
	cfg.VerzDbConfiguration = dbConfiguration
	return nil
}

func (cfg *Config) parseDialerDbConfiguration() error {
	if err := cfg.Validate(); err != nil {
		log.Println("parseDialerDbConfiguration: ", err)
		return err
	}
	if cfg.DialerDbConfigurationFile == "" {
		log.Println("empty file name:", cfg.DialerDbConfigurationFile)
		return nil
	}
	if cfg.DialerDbConfiguration != nil {
		log.Println("DialerDbConfiguration already loaded")
		return nil
	}
	dbConfiguration := &DbConfiguration{}
	jsonConfiguration, err := cfg.file.ReadConfigurationFile(cfg.DialerDbConfigurationFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonConfiguration, &dbConfiguration)
	if err != nil {
		return err
	}
	if dbConfiguration.Validate() != nil {
		return dbConfiguration.Validate()
	}
	cfg.DialerDbConfiguration = dbConfiguration
	return nil
}

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
	if err = cfg.parseDbConfiguration(); err != nil {
		return cfg, err
	}
	if err = cfg.parseZimpeDbConfiguration(); err != nil {
		return cfg, err
	}
	if err = cfg.parseVerzDbConfiguration(); err != nil {
		return cfg, err
	}
	if err = cfg.parseDialerDbConfiguration(); err != nil {
		return cfg, err
	}
	return cfg, nil
}
