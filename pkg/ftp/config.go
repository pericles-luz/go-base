package ftp

import (
	"encoding/json"

	"github.com/pericles-luz/go-base/pkg/conf"
)

type Config struct {
	file     conf.ConfigBase
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *Config) Load(file string) error {
	raw, err := b.file.ReadConfigurationFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), b)
	if err != nil {
		return err
	}
	return nil
}

func (b *Config) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"url":      b.URL,
		"username": b.Username,
		"password": b.Password,
	}
}

func NewConfig() *Config {
	return &Config{}
}
