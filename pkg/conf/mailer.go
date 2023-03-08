package conf

import "encoding/json"

type Mailer struct {
	file     ConfigBase
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (b *Mailer) Load(file string) error {
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

func (b *Mailer) GetConfig() map[string]string {
	return map[string]string{
		"host":     b.Host,
		"port":     b.Port,
		"username": b.Username,
		"password": b.Password,
	}
}

func NewMailer() *Mailer {
	return &Mailer{}
}
