package digisac

import (
	"errors"
	"log"
	"time"

	"github.com/pericles-luz/go-base/pkg/rest"
)

const (
	TOKEN_VALIDITY_MINUTES = 60 * 3
)

type Digisac struct {
	token  *rest.Token
	rest   *rest.Rest
	parser *Parser
}

func (d *Digisac) Autenticate() error {
	if d.token != nil && d.token.IsValid() {
		return nil
	}
	if d.getRest().GetConfig("token") == "" {
		return errors.New("token não encontrado")
	}
	token := rest.NewToken()
	token.SetKey(d.getRest().GetConfig("token"))
	token.SetValidity(time.Now().Add(TOKEN_VALIDITY_MINUTES * time.Minute).Format("2006-01-02 15:04:05"))
	if !token.IsValid() {
		log.Println("Validade ruim:", token.GetValidity())
		return errors.New("token inválido")
	}
	log.Println("Validade passou:", token.GetValidity())
	d.token = token
	d.getRest().SetToken(token)
	return nil
}

func (d *Digisac) getRest() *rest.Rest {
	return d.rest
}

func (d *Digisac) GetContactByID(id string) (*Contact, error) {
	resp, err := d.get(d.getRest().GetConfig("linkAPI")+"/contacts/"+id, nil)
	if err != nil {
		return nil, err
	}
	return d.getParser().GetContact(resp.GetRaw())
}

func (d *Digisac) getParser() *Parser {
	return d.parser
}

func (d *Digisac) get(url string, data map[string]interface{}) (*rest.Response, error) {
	if err := d.Autenticate(); err != nil {
		return nil, err
	}
	log.Println("data para o GET: ", data)

	resp, err := d.getRest().Get(data, url)
	if err != nil {
		return nil, err
	}
	log.Println("resposta do GET: ", resp)
	return resp, nil
}

func NewDigisac(rest *rest.Rest) *Digisac {

	d := &Digisac{}
	d.rest = rest
	d.parser = NewParser()
	return d
}
