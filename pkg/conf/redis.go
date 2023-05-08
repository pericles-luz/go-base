package conf

import "encoding/json"

type Redis struct {
	file     ConfigBase
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func (r *Redis) Load(file string) error {
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

func (r *Redis) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"DE_Host":  r.Host,
		"NU_Port":  r.Port,
		"PW_Senha": r.Password,
		"NU_Banco": r.DB,
	}
}

func NewRedis() *Redis {
	return &Redis{}
}
