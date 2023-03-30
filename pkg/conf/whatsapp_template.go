package conf

import "encoding/json"

type WhatsappTemplate struct {
	file      ConfigBase
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func (b *WhatsappTemplate) Load(file string) error {
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

func (b *WhatsappTemplate) GetConfig() map[string]string {
	return map[string]string{
		"name":      b.Name,
		"namespace": b.Namespace,
	}
}

func NewWhatsappTemplate() *WhatsappTemplate {
	return &WhatsappTemplate{}
}
