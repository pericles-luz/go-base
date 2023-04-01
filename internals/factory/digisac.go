package factory

import (
	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/pericles-luz/go-base/pkg/digisac"
	"github.com/pericles-luz/go-base/pkg/rest"
)

func NewDigisac(file string) (*digisac.Digisac, error) {
	config := conf.NewDigisac()
	err := config.Load(file)
	if err != nil {
		return nil, err
	}
	rest := rest.NewRest(config.GetConfig())
	digisac := digisac.NewDigisac(rest)
	return digisac, nil
}
