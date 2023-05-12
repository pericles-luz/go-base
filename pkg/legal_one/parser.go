package legal_one

import (
	"errors"

	"github.com/pericles-luz/go-base/pkg/utils"
)

type Parser struct {
	data map[string]interface{}
}

func (p *Parser) AuthResponse(data string) (*AuthResponse, error) {
	response := &AuthResponse{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) GetContactResponse(data string) (*ContactResponse, error) {
	response := &ContactResponse{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) GetLawsuitResponse(data string) (*LawsuitResponse, error) {
	response := &LawsuitResponse{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) GetAppealResponse(data string) (*AppealResponse, error) {
	response := &AppealResponse{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) GetLitigationResponse(data string) (*LitigationResponse, error) {
	response := &LitigationResponse{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) IndividualRegistrateRequest() (*Individual, error) {
	if err := p.validateIndividualRegistrateRequest(); err != nil {
		return nil, errors.New("invalid individual")
	}
	individual := &Individual{}
	individual.Name = p.data["DE_Pessoa"].(string)
	individual.IdentificationNumber = p.data["CO_CPF"].(string)
	return individual, nil
}

func (p *Parser) ParticipationRegistrateRequest() (*LitigationParticipant, error) {
	if err := p.validateParticipationRegistrateRequest(); err != nil {
		return nil, errors.New("invalid participation")
	}
	participation := &LitigationParticipant{}
	participation.Type = p.data["DE_Tipo"].(string)
	participation.ContactID = p.data["ID_Contato"].(int)
	participation.PositionID = p.data["ID_Posicao"].(int)
	participation.IsMainParticipant = p.data["SN_Principal"].(bool)
	return participation, nil
}

func (p *Parser) IndividualRegistrateResponse(data string) (*Individual, error) {
	response := &Individual{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) ParticipationRegistrateResponse(data string) (*LitigationParticipant, error) {
	response := &LitigationParticipant{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) LitigationParticipationResponse(data string) (*LitigationParticipationResponse, error) {
	response := &LitigationParticipationResponse{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Parser) ResponseError(data string) (string, error) {
	response := &ResponseError{}
	err := utils.ByteToStruct([]byte(data), response)
	if err != nil {
		return "", err
	}
	return response.Error.Message, nil
}

func (p *Parser) getData() map[string]interface{} {
	return p.data
}

func (p *Parser) setData(data map[string]interface{}) {
	p.data = data
}

func NewParser() *Parser {
	return &Parser{}
}
