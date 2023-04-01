package conf

import (
	"encoding/json"
)

type Digisac struct {
	file    ConfigBase
	LinkAPI string `json:"linkAPI"`
	Token   string `json:"token"`
}

func (b *Digisac) Load(file string) error {
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

func (b *Digisac) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"linkAPI": b.LinkAPI,
		"token":   b.Token,
	}
}

func NewDigisac() *Digisac {
	return &Digisac{}
}
