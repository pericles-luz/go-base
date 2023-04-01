package digisac_test

func dataWebhookMessageCreatedReceiving() string {
	return `{
		"event":"message.created",
		"data":{
			"id":"2d9d1b57-db7c-4586-9c73-91b1b571c2b4",
			"isFromMe":false,
			"sent":true,
			"type":"chat",
			"timestamp":"2023-04-01T13:06:34.100Z",
			"data":{
				"ack":-1,
				"isNew":true,
				"isFirst":false
			},
			"visible":true,
			"accountId":"152b1195-e23d-4cf8-84cf-8c8c69b6d8e6",
			"contactId":"ce4af5d4-afd5-4c4f-b5ab-a9ade3c223b0",
			"fromId":"ce4af5d4-afd5-4c4f-b5ab-a9ade3c223b0",
			"serviceId":"5f4e6542-852c-4ca4-8331-a7f92eef9121",
			"toId":null,
			"userId":null,
			"ticketId":"df9e4a0a-d02c-4af1-92d1-32f6d18ea6a6",
			"ticketUserId":null,
			"ticketDepartmentId":null,
			"quotedMessageId":null,
			"origin":null,
			"hsmId":null,
			"text":"irpf",
			"obfuscated":false,
			"files":null,
			"quotedMessage":null,
			"isFromBot":false
		},
		"webhookId":"34f58096-8e65-48af-9914-3de5f7b4cdc9",
		"timestamp":"2023-04-01T13:06:35.104Z"
	}`
}

func dataWebhookMessageCreatedTransfering() string {
	return `{
		"event":"message.created",
		"data":{
			"id":"23b89b8c-b2dc-4629-aa28-fe50d23694c9",
			"isFromMe":false,
			"sent":false,
			"type":"ticket",
			"timestamp":"2023-04-01T13:13:05.182Z",
			"data":{
				"ticketTransfer":true
			},
			"visible":true,
			"accountId":"152b1195-e23d-4cf8-84cf-8c8c69b6d8e6",
			"contactId":"ce4af5d4-afd5-4c4f-b5ab-a9ade3c223b0",
			"fromId":null,
			"serviceId":"5f4e6542-852c-4ca4-8331-a7f92eef9121",
			"toId":null,
			"userId":null,
			"ticketId":"df9e4a0a-d02c-4af1-92d1-32f6d18ea6a6",
			"ticketUserId":null,
			"ticketDepartmentId":null,
			"quotedMessageId":null,
			"origin":"ticket",
			"hsmId":null,
			"obfuscated":false,
			"files":null,
			"quotedMessage":null,
			"isFromBot":false
		},
		"webhookId":"34f58096-8e65-48af-9914-3de5f7b4cdc9",
		"timestamp":"2023-04-01T13:13:05.277Z"
	}`
}

func dataWebhookMessageCreatedSending() string {
	return `{
		"event":"message.created",
		"data":{
			"id":"710e6012-a5c1-4ddb-9ad6-e48135f3406f",
			"isFromMe":true,
			"sent":true,
			"type":"chat",
			"timestamp":"2023-04-01T13:14:55.259Z",
			"data":{
				"ack":0,
				"isNew":true,
				"isFirst":false,
				"dontOpenTicket":false
			},
			"visible":true,
			"accountId":"152b1195-e23d-4cf8-84cf-8c8c69b6d8e6",
			"contactId":"ce4af5d4-afd5-4c4f-b5ab-a9ade3c223b0",
			"fromId":null,
			"serviceId":"5f4e6542-852c-4ca4-8331-a7f92eef9121",
			"toId":null,
			"userId":"f2541200-958c-4255-a597-e350485212a2",
			"ticketId":"df9e4a0a-d02c-4af1-92d1-32f6d18ea6a6",
			"ticketUserId":null,
			"ticketDepartmentId":null,
			"quotedMessageId":null,
			"origin":"user",
			"hsmId":null,
			"text":"enviando arquivo",
			"obfuscated":false,
			"files":null,
			"quotedMessage":null,
			"isFromBot":false
		},
		"webhookId":"34f58096-8e65-48af-9914-3de5f7b4cdc9",
		"timestamp":"2023-04-01T13:14:55.318Z"
	}`
}

func dataWebhookMessageCreatedClosing() string {
	return `{
		"event":"message.created",
		"data":{
			"id":"a1d55ed1-8514-4438-b70d-404e541fbc9c",
			"isFromMe":false,
			"sent":false,
			"type":"ticket",
			"timestamp":"2023-04-01T13:16:48.600Z",
			"data":{
				"ticketClose":true
			},
			"visible":true,
			"accountId":"152b1195-e23d-4cf8-84cf-8c8c69b6d8e6",
			"contactId":"ce4af5d4-afd5-4c4f-b5ab-a9ade3c223b0",
			"fromId":null,
			"serviceId":"5f4e6542-852c-4ca4-8331-a7f92eef9121",
			"toId":null,
			"userId":null,
			"ticketId":"df9e4a0a-d02c-4af1-92d1-32f6d18ea6a6",
			"ticketUserId":null,
			"ticketDepartmentId":null,
			"quotedMessageId":null,
			"origin":"ticket",
			"hsmId":null,
			"obfuscated":false,
			"files":null,
			"quotedMessage":null,
			"isFromBot":false
		},
		"webhookId":"34f58096-8e65-48af-9914-3de5f7b4cdc9",
		"timestamp":"2023-04-01T13:16:48.689Z"
	}`
}
