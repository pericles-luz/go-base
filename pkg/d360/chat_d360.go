package d360

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/pericles-luz/go-base/pkg/rest"
	"github.com/pericles-luz/go-base/pkg/utils"
)

const (
	TOKEN_VALIDITY_MINUTES = 60 * 3
)

type D360 struct {
	parser *D360_Parser
	rest   *rest.Rest
	token  *rest.Token
}

func (d *D360) getParser() *D360_Parser {
	return d.parser
}

func (d *D360) getRest() *rest.Rest {
	return d.rest
}

func (d *D360) Autenticate() (rest.IToken, error) {
	if d.token != nil && d.token.IsValid() {
		return d.token, nil
	}
	if d.getRest().GetConfig("token") == "" {
		return nil, errors.New("token não encontrado")
	}
	token := rest.NewToken()
	token.SetKey(d.getRest().GetConfig("token"))
	token.SetValidity(time.Now().Add(TOKEN_VALIDITY_MINUTES * time.Minute).Format("2006-01-02 15:04:05"))
	if !token.IsValid() {
		return nil, errors.New("token inválido")
	}
	d.token = token
	return d.token, nil
}

func (d *D360) SendMessage(data map[string]interface{}) ([]rest.ISendMessageResponse, error) {
	d.getParser().setData(data)
	message, err := d.getParser().sendMessageResquest()
	if err != nil {
		return nil, err
	}
	requestData, err := utils.StructToMapInterface(message)
	if err != nil {
		return nil, err
	}
	resp, err := d.post(d.getRest().GetConfig("linkAPI")+"/messages", requestData)
	if err != nil {
		return nil, err
	}
	messageResponse, err := d.getParser().sendMessageResponse(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	return messageResponse, nil
}

func (d *D360) SendMessageInteractive(data map[string]interface{}) ([]rest.ISendMessageResponse, error) {
	d.getParser().setData(data)
	message, err := d.getParser().SendInteractiveMessageResquest()
	if err != nil {
		return nil, err
	}
	requestData, err := utils.StructToMapInterface(message)
	if err != nil {
		return nil, err
	}
	resp, err := d.post(d.getRest().GetConfig("linkAPI")+"/messages", requestData)
	if err != nil {
		return nil, err
	}
	messageResponse, err := d.getParser().sendMessageResponse(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	return messageResponse, nil
}

func (d *D360) SendMessageTemplate(data map[string]interface{}) ([]rest.ISendMessageResponse, error) {
	d.getParser().setData(data)
	message, err := d.getParser().SendTemplateMessage()
	if err != nil {
		return nil, err
	}
	requestData, err := utils.StructToMapInterface(message)
	if err != nil {
		return nil, err
	}
	resp, err := d.post(d.getRest().GetConfig("linkAPI")+"/messages", requestData)
	if err != nil {
		return nil, err
	}
	messageResponse, err := d.getParser().sendMessageResponse(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	return messageResponse, nil
}

func (d *D360) GetTemplateInteractive() (*D360_TemplateInteractiveResponse, error) {
	resp, err := d.get(d.getRest().GetConfig("linkAPI")+"/configs/templates", nil)
	if err != nil {
		return nil, err
	}
	messageResponse, err := d.getParser().templateInteractiveResponse(resp.GetRaw())
	if err != nil {
		return nil, err
	}
	return messageResponse, nil
}

func (d *D360) post(url string, data map[string]interface{}) (*rest.Response, error) {
	if _, err := d.Autenticate(); err != nil {
		return nil, err
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	log.Println("dataJson para o POST: ", string(dataJson))
	resp, err := d.getRest().PostWithHeaderNoAuth(data, url, map[string]string{
		"D360-API-KEY": d.token.GetKey(),
		"Content-Type": "application/json",
	})
	if err != nil {
		return nil, err
	}
	log.Println("resposta do POST: ", resp)
	// response, err := d.getParser().ResponseError(resp.GetRaw())
	// if err != nil {
	// 	return nil, err
	// }
	// if response != "" {
	// 	return nil, errors.New(response)
	// }
	return resp, nil
}

func (d *D360) get(url string, data map[string]interface{}) (*rest.Response, error) {
	if _, err := d.Autenticate(); err != nil {
		return nil, err
	}
	log.Println("data para o GET: ", data)

	resp, err := d.getRest().GetWithHeader(data, url, map[string]string{
		"D360-API-KEY": d.token.GetKey(),
	})
	if err != nil {
		return nil, err
	}
	log.Println("resposta do GET: ", resp)
	// response, err := d.getParser().ResponseError(resp.GetRaw())
	// if err != nil {
	// 	return nil, err
	// }
	// if response != "" {
	// 	return nil, errors.New(response)
	// }
	return resp, nil
}

func NewChatD360(rest *rest.Rest) *D360 {
	return &D360{
		rest:   rest,
		parser: NewD360Parser(nil),
	}
}
