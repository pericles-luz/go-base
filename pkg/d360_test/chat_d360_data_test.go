package d360_test

func dataChatD360MessageRequest() string {
	return `{
    "to": "31986058910",
    "type": "text",
    "text": {
        "body": "Teste enviado com sucesso!"
    }
    }`
}

func dataChatD360MessageContactsResponse() string {
	return `{
        "contacts":[{
            "input":"553186058910",
            "wa_id":"553186058910"
        }],
        "messages":[{
            "id":"gBEGVTGGBYkQAgmxLiuri5r5qMI"
        }],
        "meta":{
            "api_status":"stable",
            "version":"2.43.2"
        }
    }`
}

func dataChatD360InteractiveMessageRequest() string {
	return `{
        "type": "interactive",
        "to": "31986058910",
        "interactive": {
            "type": "button",
            "header": {
                "type": "text",
                "text": "Teste de header mensagem interativa"
            },
            "body": {
                "text": "Teste de body mensagem interativa"
            },
            "footer": {
                "text": "Teste de footer mensagem interativa"
            },
            "action": {
                "buttons": [{
                    "type": "reply",
                    "reply": {
                        "id": "12345-s",
                        "title": "Sim"
                    }
                },
                {
                    "type": "reply",
                    "reply": {
                        "id": "12345-n",
                        "title": "Não"
                    }
                }]
            }
        }
    }`
}

func dataChatD360InteractiveTemplateResponse() string {
	return `{
    "count": 3,
    "filters": {
    },
    "limit": 1000,
    "offset": 0,
    "sort": [
        "id"
    ],
    "total": 3,
    "waba_templates": [
        {
            "category": "ACCOUNT_UPDATE",
            "components": [
                {
                    "format": "TEXT",
                    "text": "Facebook Template Message",
                    "type": "HEADER"
                },
                {
                    "text": "Thank you for your request. \nAs Facebook closes the conversation after 24hours we cannot reach out to you anymore besides this template. \nIf you like to get in contact with us again and revoke the conversation please just click on one of the buttons and we get back to you.",
                    "type": "BODY"
                },
                {
                    "text": "Many thanks, your 360dialog team",
                    "type": "FOOTER"
                },
                {
                    "buttons": [
                        {
                            "text": "RESOLVED",
                            "type": "QUICK_REPLY"
                        },
                        {
                            "text": "NO",
                            "type": "QUICK_REPLY"
                        },
                        {
                            "text": "YES",
                            "type": "QUICK_REPLY"
                        }
                    ],
                    "type": "BUTTONS"
                }
            ],
            "language": "en",
            "name": "test_1",
            "namespace": "xxxxxxxx_xxxx_xxxx_xxxx_xxxxxxxxxxxx",
            "rejected_reason": "NONE",
            "status": "APPROVED"
        },
        {
            "category": "TICKET_UPDATE",
            "components": [
                {
                    "text": "Thank you for reaching out to us. We are looking into your request and get back to you.",
                    "type": "BODY"
                },
                {
                    "text": "Feel free to get in contact",
                    "type": "FOOTER"
                },
                {
                    "buttons": [
                        {
                            "phone_number": "+49xxxxxxxxx",
                            "text": "Call number",
                            "type": "PHONE_NUMBER"
                        },
                        {
                            "text": "360dialog Website",
                            "type": "URL",
                            "url": "https://www.360dialog.com/de/whatsapp-business/"
                        }
                    ],
                    "type": "BUTTONS"
                }
            ],
            "language": "en",
            "name": "test_2",
            "namespace": "xxxxxxxx_xxxx_xxxx_xxxx_xxxxxxxxxxxx",
            "rejected_reason": "NONE",
            "status": "APPROVED"
        },
        {
            "category": "TICKET_UPDATE",
            "components": [
                {
                    "text": "Hello {{1}}! This is a test message, setup with API version {{2}}",
                    "type": "BODY"
                },
                {
                    "buttons": [
                        {
                            "text": "Follow the link!",
                            "type": "URL",
                            "url": "https://www.360dialog.com/{{1}}"
                        }
                    ],
                    "type": "BUTTONS"
                },
                {
                    "text": "This is the footer",
                    "type": "FOOTER"
                }
            ],
            "language": "en",
            "name": "test_3",
            "namespace": "xxxxxxxx_xxxx_xxxx_xxxx_xxxxxxxxxxxx",
            "rejected_reason": "NONE",
            "status": "APPROVED"
        }
    ]
	}`
}

func dataChatD360SendMessageWithTemplateMediaAndButtons() string {
	return `{
        "to": "3186058910",
        "type": "template",
        "template": {
            "namespace": "f6d29be0_b414_48ba_b2f0_34c8753ce701",
            "language": {
                "policy": "deterministic",
                "code": "pt_BR"
            },
            "name": "simulacao_crm",
            "components":[{
                "type": "header",
                "parameters": [{
                    "type":"image",
                    "image": {
                        "link": "https://connectpeoplebrasil.com.br/assets/images/SaporeMode.png"
                    }
                }]
            }]
        }
    }`
}

func dataChatD360InteractiveTemplateResponseWithHandle() string {
	return `{
        "count": 1,
        "filters": {
        },
        "limit": 1000,
        "offset": 0,
        "sort": [
          "id"
        ],
        "total": 1,
        "waba_templates": [
          {
            "category": "TRANSACTIONAL",
            "components": [
              {
                "text": "Olá, conforme combinado, segue sua simulação. 🙂\n\nDeseja continuar com o atendimento?  ",
                "type": "BODY"
              },
              {
                "example": {
                  "header_handle": [
                    "https://connectpeoplebrasil.com.br/assets/images/SaporeMode.png"
                  ]
                },
                "format": "IMAGE",
                "type": "HEADER"
              },
              {
                "buttons": [
                  {
                    "text": "Sim",
                    "type": "QUICK_REPLY"
                  },
                  {
                    "text": "Não",
                    "type": "QUICK_REPLY"
                  }
                ],
                "type": "BUTTONS"
              }
            ],
            "language": "pt_BR",
            "name": "simulacao_crm",
            "namespace": "f6d29be0_b414_48ba_b2f0_34c8753ce701",
            "rejected_reason": "",
            "status": "approved"
          }
        ]
      }`
}

func dataChatD360InteractiveMessageMap() map[string]interface{} {
	return map[string]interface{}{
		"DE_Telefone": "31986058910",
		"interactive": map[string]interface{}{
			"DE_Tipo": "button",
			"cabecalho": map[string]interface{}{
				"DE_Tipo":  "text",
				"DE_Texto": "Pesquisa importante",
			},
			"corpo": map[string]interface{}{
				"DE_Texto": "O seu time planeja perder em casa para o maior rival",
			},
			"rodape": map[string]interface{}{
				"DE_Texto": "Você concorda?",
			},
			"acao": map[string]interface{}{
				"botoes": []map[string]interface{}{
					{
						"DE_Tipo": "reply",
						"resposta": map[string]interface{}{
							"ID_Botao":  "12345-s",
							"DE_Titulo": "Sim",
						},
					},
					{
						"DE_Tipo": "reply",
						"resposta": map[string]interface{}{
							"ID_Botao":  "12345-n",
							"DE_Titulo": "Não",
						},
					},
				},
			},
		},
	}
}

func dataChatD360TextTemplateMessageMap() map[string]interface{} {
	return map[string]interface{}{
		"DE_Telefone": "31986058910",
		"template": map[string]interface{}{
			"DE_Namespace": "39751bde_f26f_42f3_b928_aa4267759d7f",
			"DE_Nome":      "token_logon",
			"componentes": []map[string]interface{}{
				{
					"DE_Tipo": "body",
					"parametros": []map[string]interface{}{
						{
							"DE_Tipo":  "text",
							"DE_Texto": "897645",
						},
					},
				},
			},
		},
	}
}

func dataChatD360TextTemplateMessageOlaMap() map[string]interface{} {
	return map[string]interface{}{
		"DE_Telefone": "31986058910",
		"template": map[string]interface{}{
			"DE_Namespace": "39751bde_f26f_42f3_b928_aa4267759d7f",
			"DE_Nome":      "primeiro_contato",
			"componentes": []map[string]interface{}{
				{
					"DE_Tipo": "body",
					"parametros": []map[string]interface{}{
						{
							"DE_Tipo":  "text",
							"DE_Texto": "Péricles",
						},
						{
							"DE_Tipo":  "text",
							"DE_Texto": "bom dia",
						},
						{
							"DE_Tipo":  "text",
							"DE_Texto": "João",
						},
					},
				},
			},
		},
	}
}

func dataChatD360InteractiveTemplateWithImageMap() map[string]interface{} {
	return map[string]interface{}{
		"DE_Telefone": "31986058910",
		"template": map[string]interface{}{
			"DE_Namespace": "f6d29be0_b414_48ba_b2f0_34c8753ce701",
			"DE_Nome":      "simulacao_crm",
			"componentes": []map[string]interface{}{
				{
					"DE_Tipo": "header",
					"parametros": []map[string]interface{}{
						{
							"DE_Tipo": "image",
							"imagem": map[string]interface{}{
								"LN_Imagem": "https://connectpeoplebrasil.com.br/assets/images/SaporeMode.png",
							},
						},
					},
				},
			},
		},
	}
}

func dataChatD360InteractiveMessageWithImageMap() map[string]interface{} {
	return map[string]interface{}{
		"DE_Telefone": "31986058910",
		"interactive": map[string]interface{}{
			"DE_Tipo": "button",
			"cabecalho": map[string]interface{}{
				"DE_Tipo": "image",
				"imagem": map[string]interface{}{
					"LN_Imagem": "https://connectpeoplebrasil.com.br/assets/images/SaporeMode.png",
				},
			},
			"corpo": map[string]interface{}{
				"DE_Texto": "Condição exclusiva para você",
			},
			"rodape": map[string]interface{}{
				"DE_Texto": "Deseja prosseguir?",
			},
			"acao": map[string]interface{}{
				"botoes": []map[string]interface{}{
					{
						"DE_Tipo": "reply",
						"resposta": map[string]interface{}{
							"ID_Botao":  "12345-s",
							"DE_Titulo": "Sim",
						},
					},
					{
						"DE_Tipo": "reply",
						"resposta": map[string]interface{}{
							"ID_Botao":  "12345-n",
							"DE_Titulo": "Não",
						},
					},
				},
			},
		},
	}
}
