package legal_one

import (
	"errors"

	"github.com/pericles-luz/go-base/pkg/utils"
)

const (
	POSITION_CREDITOR           = 24 // exequente
	PARTICIPATION_TYPE_CUSTOMER = "Customer"
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

func (p *Parser) validateParticipationRegistrateRequest() error {
	if _, ok := p.getData()["ID_Contato"]; !ok {
		return errors.New("contactID is required")
	}
	if _, ok := p.getData()["ID_Acao"]; !ok {
		return errors.New("lawsuitID is required")
	}
	if p.getData()["ID_Posicao"] == nil {
		p.getData()["ID_Posicao"] = POSITION_CREDITOR
	}
	if p.getData()["DE_Tipo"] == nil {
		p.getData()["DE_Tipo"] = PARTICIPATION_TYPE_CUSTOMER
	}
	if p.getData()["SN_Principal"] == nil {
		p.getData()["SN_Principal"] = false
	}
	return nil
}
