package conf

import "encoding/json"

type LegalOne struct {
	file     ConfigBase
	User     string `json:"user"`
	Password string `json:"password"`
	LinkAuth string `json:"linkAuth"`
	LinkAPI  string `json:"linkAPI"`
}

func (l *LegalOne) Load(file string) error {
	raw, err := l.file.ReadConfigurationFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(raw), l)
	if err != nil {
		return err
	}
	return nil
}

func (l *LegalOne) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"DE_User":  l.User,
		"PW_Senha": l.Password,
		"LN_Auth":  l.LinkAuth,
		"LN_API":   l.LinkAPI,
	}
}

func NewLegalOne() *LegalOne {
	return &LegalOne{}
}
