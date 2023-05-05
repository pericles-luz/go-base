package legal_one

import (
	"errors"

	"github.com/pericles-luz/go-base/pkg/utils"
)

func (p *Parser) validateIndividualRegistrateRequest() error {
	if p.getData()["DE_Pessoa"] == nil {
		return errors.New("name is required")
	}
	if p.getData()["CO_CPF"] == nil {
		return errors.New("identificationNumber is required")
	}
	cpf, err := utils.FormatCPF(p.getData()["CO_CPF"].(string))
	if err != nil {
		return err
	}
	p.getData()["CO_CPF"] = cpf
	return nil
}
