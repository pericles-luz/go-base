package conf

import (
	"encoding/json"
)

type ChatD360Config struct {
	file         ConfigBase
	LinkAPI      string `json:"linkAPI"`
	LinkAuth     string `json:"linkAuth"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	GrantType    string `json:"grantType"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
	ServiceID    string `json:"serviceId"`
	ServiceType  string `json:"serviceType"`
	Token        string `json:"token"`
}

type IChatD360Config interface {
	Load(file string) error
	GetConfig() map[string]string
}

func (b *ChatD360Config) Load(file string) error {
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

func (b *ChatD360Config) GetConfig() map[string]interface{} {
	return map[string]interface{}{
		"linkAPI":      b.LinkAPI,
		"linkAuth":     b.LinkAuth,
		"clientID":     b.ClientID,
		"clientSecret": b.ClientSecret,
		"grantType":    b.GrantType,
		"username":     b.Username,
		"password":     b.Password,
		"scope":        b.Scope,
		"serviceId":    b.ServiceID,
		"serviceType":  b.ServiceType,
		"token":        b.Token,
	}
}

func NewChatD360Config() *ChatD360Config {
	return &ChatD360Config{}
}
