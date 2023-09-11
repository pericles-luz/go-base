package d360

import (
	"encoding/json"

	"github.com/pericles-luz/go-base/internal/interfaces"
)

type D360_Parser struct {
	data          map[string]interface{}
	isInteractive bool
}

func (d *D360_Parser) setData(data map[string]interface{}) {
	d.data = data
}

func (d *D360_Parser) getData() map[string]interface{} {
	return d.data
}

func (d *D360_Parser) sendMessageResquest() (*D360_MessageRequest, error) {
	err := d.validateSendMessageResquest()
	if err != nil {
		return nil, err
	}
	result := &D360_MessageRequest{}
	result.RecipientType = "individual"
	result.Type = "text"
	result.To = FormatPhonenumber(d.getData()["DE_Telefone"].(string))
	result.Text.Body = d.getData()["TX_Mensagem"].(string)
	return result, nil
}

func (d *D360_Parser) SendInteractiveMessageResquest() (*D360_MessageInteractiveRequest, error) {
	err := d.validadeMessageInteractiveRequest()
	if err != nil {
		return nil, err
	}
	result := &D360_MessageInteractiveRequest{}
	result.RecipientType = "individual"
	result.Type = "interactive"
	result.To = FormatPhonenumber(d.getData()["DE_Telefone"].(string))
	result.Interactive.Type = "button"
	// cabeçalho é opcional
	if d.getData()["interactive"].(map[string]interface{})["cabecalho"] != nil {
		result.Interactive.Header.Type = d.getData()["interactive"].(map[string]interface{})["cabecalho"].(map[string]interface{})["DE_Tipo"].(string)
		if result.Interactive.Header.Type == "text" {
			result.Interactive.Header.Text = d.getData()["interactive"].(map[string]interface{})["cabecalho"].(map[string]interface{})["DE_Texto"].(string)
		}
		if result.Interactive.Header.Type == "image" {
			image := d.getData()["interactive"].(map[string]interface{})["cabecalho"].(map[string]interface{})["imagem"].(map[string]interface{})
			result.Interactive.Header.Image.Link = image["LN_Imagem"].(string)
			// legenda é opcional
			if image["DE_Texto"] != nil {
				result.Interactive.Header.Image.Text = image["DE_Texto"].(string)
			}
		}
		if result.Interactive.Header.Type == "document" {
			document := d.getData()["interactive"].(map[string]interface{})["cabecalho"].(map[string]interface{})["documento"].(map[string]interface{})
			result.Interactive.Header.Document.Link = document["LN_Documento"].(string)
			// nome do arquivo é opcional
			if document["DE_Documento"] != nil {
				result.Interactive.Header.Document.Filename = document["DE_Documento"].(string)
			}
		}
	}
	// corpo é obrigatório
	result.Interactive.Body.Text = d.getData()["interactive"].(map[string]interface{})["corpo"].(map[string]interface{})["DE_Texto"].(string)
	// rodapé é opcional
	if d.getData()["interactive"].(map[string]interface{})["rodape"] != nil {
		result.Interactive.Footer.Text = d.getData()["interactive"].(map[string]interface{})["rodape"].(map[string]interface{})["DE_Texto"].(string)
	}
	// ação é obrigatória
	action := D360_Action{}
	for _, v := range d.getData()["interactive"].(map[string]interface{})["acao"].(map[string]interface{})["botoes"].([]map[string]interface{}) {
		button := D360_Button{}
		button.Type = v["DE_Tipo"].(string)
		if button.Type == "reply" {
			button.Reply.ID = v["resposta"].(map[string]interface{})["ID_Botao"].(string)
			button.Reply.Title = v["resposta"].(map[string]interface{})["DE_Titulo"].(string)
		}
		action.Buttons = append(action.Buttons, button)
	}
	result.Interactive.Action = action
	return result, nil
}

func (d *D360_Parser) SendTemplateMessage() (*D360_MessageTemplateRequest, error) {
	err := d.validadeMessageTemplateRequest()
	if err != nil {
		return nil, err
	}
	result := &D360_MessageTemplateRequest{}
	result.Type = "template"
	result.To = FormatPhonenumber(d.getData()["DE_Telefone"].(string))
	template := d.getData()["template"].(map[string]interface{})
	result.Template.Namespace = template["DE_Namespace"].(string)
	result.Template.Name = template["DE_Nome"].(string)
	result.Template.Language.Code = "pt_BR"
	result.Template.Language.Policy = "deterministic"
	for _, v := range template["componentes"].([]map[string]interface{}) {
		component := D360_TemplateComponent{}
		component.Type = v["DE_Tipo"].(string)
		if component.Type == "header" {
			parameters := v["parametros"].([]map[string]interface{})
			for _, p := range parameters {
				parameter := D360_TemplateParameter{}
				parameter.Type = p["DE_Tipo"].(string)
				if parameter.Type == "image" {
					image := p["imagem"].(map[string]interface{})
					parameter.Image.Link = image["LN_Imagem"].(string)
					if image["DE_Texto"] != nil {
						parameter.Image.Text = image["DE_Texto"].(string)
					}
				}
				component.Parameters = append(component.Parameters, parameter)
			}
		}
		if component.Type == "body" {
			parameters := v["parametros"].([]map[string]interface{})
			for _, p := range parameters {
				parameter := D360_TemplateParameter{}
				parameter.Type = p["DE_Tipo"].(string)
				if parameter.Type == "image" {
					image := p["imagem"].(map[string]interface{})
					parameter.Image.Link = image["LN_Imagem"].(string)
					if image["DE_Texto"] != nil {
						parameter.Image.Text = image["DE_Texto"].(string)
					}
				}
				if parameter.Type == "text" {
					parameter.Text = p["DE_Texto"].(string)
				}
				component.Parameters = append(component.Parameters, parameter)
			}
		}
		result.Template.Components = append(result.Template.Components, component)
	}
	return result, nil
}

func FormatPhonenumber(phoneNumber string) string {
	// retira o terceiro dígito do número
	return "55" + phoneNumber[0:2] + phoneNumber[3:]
}

func (d *D360_Parser) sendMessageResponse(data string) ([]interfaces.ISendMessageResponse, error) {
	response := &D360_MessageResponse{}
	err := json.Unmarshal([]byte(data), response)
	if err != nil {
		return nil, err
	}
	result := make([]interfaces.ISendMessageResponse, len(response.Messages))
	for i, v := range response.Messages {
		result[i] = &v
	}
	return result, nil
}

func (d *D360_Parser) templateInteractiveResponse(raw string) (*D360_TemplateInteractiveResponse, error) {
	response := &D360_TemplateInteractiveResponse{}
	err := json.Unmarshal([]byte(raw), response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (d *D360_Message) GetID() string {
	return d.ID
}

func NewD360Parser(data map[string]interface{}) *D360_Parser {
	return &D360_Parser{
		data: data,
	}
}
