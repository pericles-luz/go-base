package legal_one

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/pericles-luz/go-base/pkg/rest"
	"github.com/pericles-luz/go-base/pkg/utils"
)

const (
	TOKEN_VALIDITY = 10
)

type LegalOne struct {
	token  *rest.Token
	rest   *rest.Rest
	parser *Parser
}

func (l *LegalOne) Autenticate() error {
	if l.token != nil && l.token.IsValid() {
		return nil
	}
	authPreBase64 := l.getRest().GetConfig("DE_User") + ":" + l.getRest().GetConfig("PW_Senha")
	authBase64 := base64.StdEncoding.EncodeToString([]byte(authPreBase64))
	resp, err := l.getRest().PostWithHeaderNoAuth(nil, l.getRest().GetConfig("LN_Auth"), map[string]string{
		"Authorization": "Basic " + authBase64,
	})
	if err != nil {
		return err
	}
	response, err := l.getParser().AuthResponse(resp.GetRaw())
	if err != nil {
		return err
	}
	token := rest.NewToken()
	token.SetKey(response.AccessToken)
	token.SetValidity(time.Now().UTC().Add(time.Minute * TOKEN_VALIDITY).Format("2006-01-02 15:04:05"))
	l.token = token
	return nil
}

func (l *LegalOne) GetContactByCPF(cpf string) (*ContactResponse, error) {
	cpf, err := utils.FormatCPF(cpf)
	if err != nil {
		return nil, err
	}
	resp, err := l.get(l.getRest().GetConfig("LN_API")+"/contacts?$filter=identificationNumber eq '"+cpf+"'", nil)
	if err != nil {
		return nil, err
	}
	return l.getParser().GetContactResponse(resp.GetRaw())
}

func (l *LegalOne) getParser() *Parser {
	return l.parser
}

func (l *LegalOne) getRest() *rest.Rest {
	return l.rest
}

func (l *LegalOne) get(url string, data map[string]interface{}) (*rest.Response, error) {
	if err := l.Autenticate(); err != nil {
		return nil, err
	}
	log.Println("data para o GET: ", data)

	l.getRest().SetToken(l.token)
	resp, err := l.getRest().Get(data, url)
	if err != nil {
		return nil, err
	}
	log.Println("resposta do GET: ", resp)
	return resp, nil
}
func NewLegalOne(rest *rest.Rest) *LegalOne {
	return &LegalOne{
		rest:   rest,
		parser: NewParser(),
	}
}
