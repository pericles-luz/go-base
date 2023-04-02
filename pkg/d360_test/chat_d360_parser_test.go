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
