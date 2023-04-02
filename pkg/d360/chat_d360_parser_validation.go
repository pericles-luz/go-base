package d360

import (
	"errors"

	"github.com/pericles-luz/go-base/pkg/utils"
)

const (
	MISSING_RESPONSE_DATA = "faltam dados para o envio da mensagem: "
)

func (d *D360_Parser) validadeMessageInteractiveRequest() error {
	d.isInteractive = true
	if d.getData()["DE_Telefone"] == nil || utils.ValidateCellphoneNumber(d.getData()["DE_Telefone"].(string)) != nil {
		return errors.New(MISSING_RESPONSE_DATA + "Telefone")
	}
	if d.getData()["interactive"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "interactive")
	}
	if err := d.validateHeader(); err != nil {
		return err
	}
	if err := d.validateBody(); err != nil {
		return err
	}
	if err := d.validateFooter(); err != nil {
		return err
	}
	if err := d.validateAction(); err != nil {
		return err
	}
	return nil
}

func (d *D360_Parser) validadeMessageTemplateRequest() error {
	if d.getData()["DE_Telefone"] == nil || utils.ValidateCellphoneNumber(d.getData()["DE_Telefone"].(string)) != nil {
		return errors.New(MISSING_RESPONSE_DATA + "Telefone")
	}
	if d.getData()["template"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "template")
	}
	template := d.getData()["template"].(map[string]interface{})
	if template["DE_Namespace"] == nil || template["DE_Namespace"].(string) == "" {
		return errors.New(MISSING_RESPONSE_DATA + "Namespace")
	}
	if template["DE_Nome"] == nil || template["DE_Nome"].(string) == "" {
		return errors.New(MISSING_RESPONSE_DATA + "Nome")
	}
	for _, component := range template["componentes"].([]map[string]interface{}) {
		if component["DE_Tipo"] == nil {
			return errors.New(MISSING_RESPONSE_DATA + "Tipo do componente")
		}
		if component["DE_Tipo"].(string) == "header" {
			for _, param := range component["parametros"].([]map[string]interface{}) {
				if param["DE_Tipo"] != nil && param["DE_Tipo"].(string) == "image" {
					image := param["imagem"].(map[string]interface{})
					if image["LN_Imagem"] == nil || !utils.ValidateURL(image["LN_Imagem"].(string)) {
						return errors.New(MISSING_RESPONSE_DATA + "Imagem do cabeçalho")
					}
				}
				if param["DE_Tipo"] != nil && param["DE_Tipo"].(string) == "text" {
					if param["DE_Texto"] == nil || len(param["DE_Texto"].(string)) == 0 {
						return errors.New(MISSING_RESPONSE_DATA + "texto do parâmento")
					}
				}
			}
		}
		if component["DE_Tipo"].(string) == "body" {
			for _, param := range component["parametros"].([]map[string]interface{}) {
				if param["DE_Tipo"] != nil && param["DE_Tipo"].(string) == "text" {
					if param["DE_Texto"] == nil || len(param["DE_Texto"].(string)) == 0 {
						return errors.New(MISSING_RESPONSE_DATA + "texto do parâmento")
					}
				}
			}
		}
	}
	return nil
}

func (d *D360_Parser) validateSendMessageResquest() error {
	if d.getData()["DE_Telefone"] == nil || utils.ValidateCellphoneNumber(d.getData()["DE_Telefone"].(string)) != nil {
		return errors.New(MISSING_RESPONSE_DATA + "Telefone")
	}
	if d.getData()["TX_Mensagem"] == nil || d.getData()["TX_Mensagem"].(string) == "" {
		return errors.New(MISSING_RESPONSE_DATA + "Mensagem")
	}
	return nil
}

func (d *D360_Parser) validateBody() error {
	if d.getData()["interactive"].(map[string]interface{})["corpo"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Corpo da mensagem")
	}
	body := d.getData()["interactive"].(map[string]interface{})["corpo"].(map[string]interface{})
	if body["DE_Tipo"] == nil || body["DE_Tipo"].(string) == "text" {
		if body["DE_Texto"] == nil {
			return errors.New(MISSING_RESPONSE_DATA + "Texto do corpo")
		}
		if body["DE_Texto"].(string) == "" {
			return errors.New(MISSING_RESPONSE_DATA + "Texto do corpo")
		}
		if len(body["DE_Texto"].(string)) > 1024 {
			return errors.New("o texto do corpo não pode ter mais de 1024 caracteres")
		}
		return nil
	}
	return errors.New("tipo de corpo da mensagem não identificado")
}

func (d *D360_Parser) validateHeader() error {
	if d.getData()["interactive"].(map[string]interface{})["cabecalho"] == nil {
		return nil
	}
	header := d.getData()["interactive"].(map[string]interface{})["cabecalho"].(map[string]interface{})
	if header["DE_Tipo"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Tipo de cabeçalho")
	}
	if header["DE_Tipo"].(string) == "text" {
		if header["DE_Texto"] == nil {
			return errors.New(MISSING_RESPONSE_DATA + "Texto do cabeçalho")
		}
		if header["DE_Texto"].(string) == "" {
			return errors.New(MISSING_RESPONSE_DATA + "Texto do cabeçalho")
		}
		if len(header["DE_Texto"].(string)) > 60 {
			return errors.New("o texto do cabeçalho não pode ter mais de 60 caracteres")
		}
		return nil
	}
	if header["DE_Tipo"].(string) == "image" {
		if err := d.validateImage(header); err != nil {
			return err
		}
		return nil
	}
	if header["DE_Tipo"].(string) == "document" {
		if err := d.validateDocument(header); err != nil {
			return err
		}
		return nil
	}
	return errors.New("tipo de cabeçalho não identificado")
}

func (d *D360_Parser) validateImage(header map[string]interface{}) error {
	if header["imagem"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Imagem")
	}
	image := header["imagem"].(map[string]interface{})
	if image["LN_Imagem"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Link da imagem")
	}
	if image["LN_Imagem"].(string) == "" {
		return errors.New(MISSING_RESPONSE_DATA + "Link da imagem")
	}
	if image["LN_Imagem"].(string) != "" && !utils.ValidateURL(image["LN_Imagem"].(string)) {
		return errors.New("link da imagem inválido")
	}
	if d.isInteractive && image["DE_Texto"] != nil {
		return errors.New("texto do cabeçalho não pode ser enviado em uma mensagem interativa")
	}
	if image["DE_Texto"] != nil && image["DE_Texto"].(string) != "" && len(image["DE_Texto"].(string)) > 1024 {
		return errors.New("o texto do cabeçalho não pode ter mais de 1024 caracteres")
	}
	return nil
}

func (d *D360_Parser) validateDocument(header map[string]interface{}) error {
	if header["documento"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Documento")
	}
	document := header["documento"].(map[string]interface{})
	if document["LN_Documento"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Link do documento")
	}
	if document["LN_Documento"].(string) == "" {
		return errors.New(MISSING_RESPONSE_DATA + "Link do documento")
	}
	if document["LN_Documento"].(string) != "" && !utils.ValidateURL(document["LN_Documento"].(string)) {
		return errors.New("link do documento inválido")
	}
	if document["DE_Documento"] != nil && document["DE_Documento"].(string) != "" && len(document["DE_Documento"].(string)) > 250 {
		return errors.New("o nome do arquivo não pode ter mais de 250 caracteres")
	}
	return nil
}

func (d *D360_Parser) validateFooter() error {
	if d.getData()["interactive"].(map[string]interface{})["rodape"] == nil {
		return nil
	}
	footer := d.getData()["interactive"].(map[string]interface{})["rodape"].(map[string]interface{})
	if footer["DE_Tipo"] == nil || footer["DE_Tipo"].(string) == "text" {
		if footer["DE_Texto"] == nil {
			return errors.New(MISSING_RESPONSE_DATA + "Texto do rodapé")
		}
		if footer["DE_Texto"].(string) == "" {
			return errors.New(MISSING_RESPONSE_DATA + "Texto do rodapé")
		}
		if len(footer["DE_Texto"].(string)) > 60 {
			return errors.New("o texto do rodapé não pode ter mais de 60 caracteres")
		}
		return nil
	}
	return errors.New("tipo de rodapé não identificado")
}

func (d *D360_Parser) validateAction() error {
	if d.getData()["interactive"].(map[string]interface{})["acao"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Ação")
	}
	if d.getData()["interactive"].(map[string]interface{})["acao"].(map[string]interface{})["botoes"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Botões")
	}
	action := d.getData()["interactive"].(map[string]interface{})["acao"].(map[string]interface{})["botoes"].([]map[string]interface{})
	if len(action) == 0 {
		return errors.New(MISSING_RESPONSE_DATA + "Botões")
	}
	if len(action) > 3 {
		return errors.New("o número máximo de botões é 3")
	}
	for _, button := range action {
		if err := d.validateButton(button); err != nil {
			return err
		}
	}
	return nil
}

func (d *D360_Parser) validateButton(button map[string]interface{}) error {
	if button["DE_Tipo"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "Tipo do botão")
	}
	if button["DE_Tipo"].(string) == "reply" {
		if button["resposta"] == nil {
			return errors.New(MISSING_RESPONSE_DATA + "Botão de resposta")
		}
		return d.validateReplyButton(button["resposta"].(map[string]interface{}))
	}
	if button["DE_Tipo"].(string) == "text" {
		if button["DE_Titulo"] == nil {
			return errors.New(MISSING_RESPONSE_DATA + "Texto do botão")
		}
		return nil
	}
	return errors.New("tipo de botão não identificado")
}

func (d *D360_Parser) validateReplyButton(button map[string]interface{}) error {
	if button["DE_Titulo"].(string) == "" {
		return errors.New(MISSING_RESPONSE_DATA + "Texto do botão de resposta")
	}
	if len(button["DE_Titulo"].(string)) > 24 {
		return errors.New("o texto do botão de resposta não pode ter mais de 24 caracteres")
	}
	if button["ID_Botao"] == nil {
		return errors.New(MISSING_RESPONSE_DATA + "ID do botão de resposta")
	}
	if len(button["ID_Botao"].(string)) > 36 {
		return errors.New("o ID do botão de resposta não pode ter mais de 36 caracteres")
	}
	return nil
}
