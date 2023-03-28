package conf

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/pericles-luz/go-base/pkg/conf"
)

// Config for the environment, read from a json file
type Config struct {
	file                    ConfigBase
	Debug                   bool
	Production              bool
	DBConfigurationFile     string
	AgnuDBConfigurationFile string
	DBConfiguration         *conf.Database
	AgnuDBConfiguration     *conf.Database
	JwtSecret               string
}

func (cfg *Config) Validate() error {
	if cfg.DBConfigurationFile == "" {
		log.Println("empty file name:", cfg.DBConfigurationFile)
		return errors.New("no db configuration defined")
	}
	return nil
}

func (cfg *Config) parseDBConfiguration() error {
	if err := cfg.Validate(); err != nil {
		return err
	}
	if cfg.DBConfiguration != nil {
		log.Println("DBConfiguration already loaded")
		return nil
	}
	dBConfiguration, err := conf.NewDatabase(cfg.DBConfigurationFile)
	if err != nil {
		return err
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
	dBConfiguration, err := conf.NewDatabase(cfg.AgnuDBConfigurationFile)
	if err != nil {
		return err
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
