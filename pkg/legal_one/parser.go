package legal_one

import "github.com/pericles-luz/go-base/pkg/utils"

type Parser struct {
	// data map[string]interface{}
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

// func (p *Parser) getData() map[string]interface{} {
// 	return p.data
// }

func NewParser() *Parser {
	return &Parser{}
}
