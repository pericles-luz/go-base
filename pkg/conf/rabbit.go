package conf

import "encoding/json"

type Rabbit struct {
	file ConfigBase
	DSN  string `json:"dsn"`
}

func (r *Rabbit) Load(file string) error {
	raw, err := r.file.ReadConfigurationFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Rabbit) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"DE_DSN": r.DSN,
	}
}

func NewRabbit() *Rabbit {
	return &Rabbit{}
}
