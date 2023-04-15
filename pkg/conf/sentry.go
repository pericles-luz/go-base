package conf

import "encoding/json"

type Sentry struct {
	file ConfigBase
	DSN  string `json:"dsn"`
}

func (b *Sentry) Load(file string) error {
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

func (b *Sentry) GetConfig() map[string]string {
	return map[string]string{
		"dsn": b.DSN,
	}
}

func NewSentry() *Sentry {
	return &Sentry{}
}
