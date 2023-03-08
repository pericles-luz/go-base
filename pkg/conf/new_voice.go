package conf

import "encoding/json"

type NewVoiceConfig struct {
	file    ConfigBase
	LinkAPI string `json:"linkAPI"`
	Code    string `json:"code"`
	Account string `json:"account"`
}

type INewVoiceConfig interface {
	Load(file string) error
	GetConfig() map[string]string
}

func (b *NewVoiceConfig) Load(file string) error {
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

func (b *NewVoiceConfig) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"linkAPI": b.LinkAPI,
		"code":    b.Code,
		"account": b.Account,
	}
}

func NewNewVoiceConfig() *NewVoiceConfig {
	return &NewVoiceConfig{}
}
