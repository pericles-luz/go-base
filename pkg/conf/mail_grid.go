package conf

import "encoding/json"

type MailGrid struct {
	file      ConfigBase
	Host      string `json:"host_smtp"`
	LinkAPI   string `json:"linkAPI"`
	Username  string `json:"usuario_smtp"`
	Password  string `json:"senha_smtp"`
	Remetente string `json:"emailRemetente"`
}

func (b *MailGrid) Load(file string) error {
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

func (b *MailGrid) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"DE_HostSMTP":    b.Host,
		"EM_Remetente":   b.Remetente,
		"DE_UsuarioSMTP": b.Username,
		"PW_UsuarioSMTP": b.Password,
		"LN_API":         b.LinkAPI,
	}
}

func NewMailGrid() *MailGrid {
	return &MailGrid{}
}
