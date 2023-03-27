package factory

import (
	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/pericles-luz/go-base/pkg/d360"
	"github.com/pericles-luz/go-base/pkg/rest"
)

func NewChatD360(configFile string) (*d360.D360, error) {
	config := conf.NewChatD360Config()
	err := config.Load(configFile)
	if err != nil {
		return nil, err
	}
	rest := rest.NewRest(config.GetConfig())
	chatD360 := d360.NewChatD360(rest)
	return chatD360, nil
}
