package digisac

import "github.com/pericles-luz/go-base/pkg/utils"

type Parser struct {
	// data          map[string]interface{}
}

func (d *Parser) GetContact(raw string) (*Contact, error) {
	result := &Contact{}
	err := utils.ByteToStruct([]byte(raw), result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func NewParser() *Parser {
	return &Parser{}
}
