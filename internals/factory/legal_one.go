package factory

import (
	"github.com/pericles-luz/go-base/pkg/conf"
	"github.com/pericles-luz/go-base/pkg/legal_one"
	"github.com/pericles-luz/go-base/pkg/rest"
)

func NewLegalOne(file string) (*legal_one.LegalOne, error) {
	config := conf.NewLegalOne()
	err := config.Load(file)
	if err != nil {
		return nil, err
	}
	rest := rest.NewRest(config.GetConfig())
	legalOne := legal_one.NewLegalOne(rest)
	return legalOne, nil
}
