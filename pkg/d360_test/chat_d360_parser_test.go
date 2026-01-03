package d360_test

import (
	"encoding/json"
	"testing"

	"github.com/pericles-luz/go-base/pkg/d360"
	"github.com/stretchr/testify/require"
)

func TestChatD360_UnMarshalToMessageRequest(t *testing.T) {
	var sendMessageResponse d360.D360_MessageResponse
	require.NoError(t, json.Unmarshal([]byte(dataChatD360MessageRequest()), &sendMessageResponse))
	require.NoError(t, json.Unmarshal([]byte(dataChatD360MessageContactsResponse()), &sendMessageResponse))
}

func TestChatD360_UnmarshalInteractiveMessageRequest(t *testing.T) {
	var interactiveRequest d360.D360_MessageInteractiveRequest
	require.NoError(t, json.Unmarshal([]byte(dataChatD360InteractiveMessageRequest()), &interactiveRequest))
}

func TestChatD360_UnmarshalInteractiveTemplateResponse(t *testing.T) {
	var templateResponse d360.D360_TemplateInteractiveResponse
	require.NoError(t, json.Unmarshal([]byte(dataChatD360InteractiveTemplateResponse()), &templateResponse))
	require.NoError(t, json.Unmarshal([]byte(dataChatD360InteractiveTemplateResponseWithHandle()), &templateResponse))
	var templateRequest d360.D360_MessageTemplateRequest
	require.NoError(t, json.Unmarshal([]byte(dataChatD360SendMessageWithTemplateMediaAndButtons()), &templateRequest))
}

func TestChatD360_SendInteractiveTemplateResquestMap(t *testing.T) {
	data := dataChatD360InteractiveTemplateWithImageMap()
	parser := d360.NewD360Parser(data)
	message, err := parser.SendTemplateMessage()
	require.NoError(t, err)
	require.Equal(t, d360.FormatPhonenumber(data["DE_Telefone"].(string)), message.To)
	// gera json
	json, err := json.Marshal(message)
	require.NoError(t, err)
	t.Log(string(json))
}

func TestChatD360_SendInteractiveMessageResquest(t *testing.T) {
	data := dataChatD360InteractiveMessageMap()
	parser := d360.NewD360Parser(data)
	message, err := parser.SendInteractiveMessageResquest()
	require.NoError(t, err)
	require.Equal(t, d360.FormatPhonenumber(data["DE_Telefone"].(string)), message.To)
	interactive := data["interactive"].(map[string]interface{})
	require.Equal(t, interactive["corpo"].(map[string]interface{})["DE_Texto"].(string), message.Interactive.Body.Text)
	require.Equal(t, interactive["rodape"].(map[string]interface{})["DE_Texto"].(string), message.Interactive.Footer.Text)
	require.Equal(t, interactive["cabecalho"].(map[string]interface{})["DE_Tipo"].(string), message.Interactive.Header.Type)
	require.Equal(t, interactive["cabecalho"].(map[string]interface{})["DE_Texto"].(string), message.Interactive.Header.Text)
	acao := interactive["acao"].(map[string]interface{})
	botoes := acao["botoes"].([]map[string]interface{})
	for i, botao := range botoes {
		require.Equal(t, botao["DE_Tipo"].(string), message.Interactive.Action.Buttons[i].Type)
		require.Equal(t, botao["resposta"].(map[string]interface{})["DE_Titulo"].(string), message.Interactive.Action.Buttons[i].Reply.Title)
		require.Equal(t, botao["resposta"].(map[string]interface{})["ID_Botao"].(string), message.Interactive.Action.Buttons[i].Reply.ID)
	}
}

func TestChatD360_SendInteractiveMessageWithImageResquest(t *testing.T) {
	data := dataChatD360InteractiveMessageWithImageMap()
	parser := d360.NewD360Parser(data)
	message, err := parser.SendInteractiveMessageResquest()
	require.NoError(t, err)
	require.Equal(t, d360.FormatPhonenumber(data["DE_Telefone"].(string)), message.To)
	interactive := data["interactive"].(map[string]interface{})
	require.Equal(t, interactive["corpo"].(map[string]interface{})["DE_Texto"].(string), message.Interactive.Body.Text)
	require.Equal(t, interactive["rodape"].(map[string]interface{})["DE_Texto"].(string), message.Interactive.Footer.Text)
	cabecalho := interactive["cabecalho"].(map[string]interface{})
	require.Equal(t, cabecalho["DE_Tipo"].(string), message.Interactive.Header.Type)
	imagem := cabecalho["imagem"].(map[string]interface{})
	if imagem["DE_Texto"] != nil {
		require.Equal(t, imagem["DE_Texto"].(string), message.Interactive.Header.Image.Text)
	}
	require.Equal(t, imagem["LN_Imagem"].(string), message.Interactive.Header.Image.Link)
	acao := interactive["acao"].(map[string]interface{})
	botoes := acao["botoes"].([]map[string]interface{})
	for i, botao := range botoes {
		require.Equal(t, botao["DE_Tipo"].(string), message.Interactive.Action.Buttons[i].Type)
		require.Equal(t, botao["resposta"].(map[string]interface{})["DE_Titulo"].(string), message.Interactive.Action.Buttons[i].Reply.Title)
		require.Equal(t, botao["resposta"].(map[string]interface{})["ID_Botao"].(string), message.Interactive.Action.Buttons[i].Reply.ID)
	}
}

func TestChatD360_SendInteractiveMessageWithImageResquestMustFailIfHasCaption(t *testing.T) {
	data := dataChatD360InteractiveMessageWithImageMap()
	data["interactive"].(map[string]interface{})["cabecalho"].(map[string]interface{})["imagem"].(map[string]interface{})["DE_Texto"] = "caption"
	parser := d360.NewD360Parser(data)
	_, err := parser.SendInteractiveMessageResquest()
	require.NotNil(t, err)
}

func TestChatD360_SendInteractiveMessageWithDocumentResquest(t *testing.T) {
	data := dataChatD360InteractiveMessageWithPDFMap()
	parser := d360.NewD360Parser(data)
	message, err := parser.SendInteractiveMessageResquest()
	require.NoError(t, err)
	require.Equal(t, d360.FormatPhonenumber(data["DE_Telefone"].(string)), message.To)
	interactive := data["interactive"].(map[string]interface{})
	require.Equal(t, interactive["corpo"].(map[string]interface{})["DE_Texto"].(string), message.Interactive.Body.Text)
	require.Equal(t, interactive["rodape"].(map[string]interface{})["DE_Texto"].(string), message.Interactive.Footer.Text)
	cabecalho := interactive["cabecalho"].(map[string]interface{})
	require.Equal(t, cabecalho["DE_Tipo"].(string), message.Interactive.Header.Type)
	document := cabecalho["documento"].(map[string]interface{})
	if document["DE_Texto"] != nil {
		require.Equal(t, document["DE_Documento"].(string), message.Interactive.Header.Document.Filename)
	}
	require.Equal(t, document["LN_Documento"].(string), message.Interactive.Header.Document.Link)
	acao := interactive["acao"].(map[string]interface{})
	botoes := acao["botoes"].([]map[string]interface{})
	for i, botao := range botoes {
		require.Equal(t, botao["DE_Tipo"].(string), message.Interactive.Action.Buttons[i].Type)
		require.Equal(t, botao["resposta"].(map[string]interface{})["DE_Titulo"].(string), message.Interactive.Action.Buttons[i].Reply.Title)
		require.Equal(t, botao["resposta"].(map[string]interface{})["ID_Botao"].(string), message.Interactive.Action.Buttons[i].Reply.ID)
	}
}

func TestChatD360_FormatPhonenumber(t *testing.T) {
	formated := d360.FormatPhonenumber("31978675897")
	require.Equal(t, "553178675897", formated)
}

func TestChatD360_ParseTemplateProvaDeVida(t *testing.T) {
	data := dataChatD360TemplateProvaDeVidaMap()
	parser := d360.NewD360Parser(data)
	message, err := parser.SendTemplateMessage()
	require.NoError(t, err)
	require.Equal(t, d360.FormatPhonenumber(data["DE_Telefone"].(string)), message.To)
	require.Equal(t, "template", message.Type)
	require.Equal(t, "whatsapp", message.MessagingProduct)

	template := data["template"].(map[string]interface{})
	require.Equal(t, template["DE_Namespace"].(string), message.Template.Namespace)
	require.Equal(t, template["DE_Nome"].(string), message.Template.Name)
	require.Equal(t, "pt_BR", message.Template.Language.Code)
	require.Equal(t, "deterministic", message.Template.Language.Policy)

	// Valida componente header
	require.Equal(t, 4, len(message.Template.Components))
	headerComponent := message.Template.Components[0]
	require.Equal(t, "header", headerComponent.Type)
	require.Equal(t, 1, len(headerComponent.Parameters))
	require.Equal(t, "image", headerComponent.Parameters[0].Type)
	require.Equal(t, "https://api.sindireceita.org.br/html/statics/assets/images/provaDeVida.png", headerComponent.Parameters[0].Image.Link)

	// Valida componente body
	bodyComponent := message.Template.Components[1]
	require.Equal(t, "body", bodyComponent.Type)
	require.Equal(t, 2, len(bodyComponent.Parameters))
	require.Equal(t, "text", bodyComponent.Parameters[0].Type)
	require.Equal(t, "Péricles", bodyComponent.Parameters[0].Text)
	require.Equal(t, "text", bodyComponent.Parameters[1].Type)
	require.Equal(t, "02/02/2026", bodyComponent.Parameters[1].Text)

	// Valida primeiro botão
	button1Component := message.Template.Components[2]
	require.Equal(t, "button", button1Component.Type)
	require.Equal(t, "URL", button1Component.SubType)
	require.Equal(t, 0, button1Component.Index)
	require.Equal(t, 1, len(button1Component.Parameters))
	require.Equal(t, "text", button1Component.Parameters[0].Type)
	require.Equal(t, "a1b2c3d4-e5f6-7890-abcd-ef1234567890", button1Component.Parameters[0].Text)

	// Valida segundo botão
	button2Component := message.Template.Components[3]
	require.Equal(t, "button", button2Component.Type)
	require.Equal(t, "URL", button2Component.SubType)
	require.Equal(t, 1, button2Component.Index)
	require.Equal(t, 1, len(button2Component.Parameters))
	require.Equal(t, "text", button2Component.Parameters[0].Type)
	require.Equal(t, "f9e8d7c6-b5a4-3210-9876-543210fedcba", button2Component.Parameters[0].Text)

	// Gera JSON para debug
	json, err := json.Marshal(message)
	require.NoError(t, err)
	t.Log(string(json))
}

func TestChatD360_NewChatD360TemplateImageMessage(t *testing.T) {
	data := d360.NewChatD360TemplateImageMessage(map[string]interface{}{
		"DE_Telefone":  "31986058910",
		"DE_Namespace": "test_namespace",
		"DE_Nome":      "test_template",
		"LN_Imagem":    "https://example.com/image.png",
		"parametros": []map[string]interface{}{
			{
				"DE_Tipo":  "text",
				"DE_Texto": "Nome Teste",
			},
		},
		"botoes": []map[string]interface{}{
			{
				"DE_Texto": "uuid-button-1",
			},
		},
	})

	require.Equal(t, "31986058910", data["DE_Telefone"])
	template := data["template"].(map[string]interface{})
	require.Equal(t, "test_namespace", template["DE_Namespace"])
	require.Equal(t, "test_template", template["DE_Nome"])

	componentes := template["componentes"].([]map[string]interface{})
	require.Equal(t, 3, len(componentes))

	// Valida header
	require.Equal(t, "header", componentes[0]["DE_Tipo"])
	headerParams := componentes[0]["parametros"].([]map[string]interface{})
	require.Equal(t, "image", headerParams[0]["DE_Tipo"])
	require.Equal(t, "https://example.com/image.png", headerParams[0]["imagem"].(map[string]interface{})["LN_Imagem"])

	// Valida body
	require.Equal(t, "body", componentes[1]["DE_Tipo"])
	bodyParams := componentes[1]["parametros"].([]map[string]interface{})
	require.Equal(t, "text", bodyParams[0]["DE_Tipo"])
	require.Equal(t, "Nome Teste", bodyParams[0]["DE_Texto"])

	// Valida botão
	require.Equal(t, "button", componentes[2]["DE_Tipo"])
	require.Equal(t, "URL", componentes[2]["DE_SubTipo"])
	require.Equal(t, 0, componentes[2]["NU_Indice"])
	buttonParams := componentes[2]["parametros"].([]map[string]interface{})
	require.Equal(t, "text", buttonParams[0]["DE_Tipo"])
	require.Equal(t, "uuid-button-1", buttonParams[0]["DE_Texto"])
}
