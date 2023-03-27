package d360

type D360_Text struct {
	Body string `json:"body"`
}

type D360_MessageRequest struct {
	To            string                 `json:"to"`
	Type          string                 `json:"type"`
	RecipientType string                 `json:"recipient_type,omitempty"`
	Text          D360_Text              `json:"text,omitempty"`
	Template      D360_TemplateToMessage `json:"template,omitempty"`
}

type D360_Message struct {
	ID string `json:"id"`
}

type D360_Contact struct {
	Input string `json:"input"`
	WaID  string `json:"wa_id"`
}

type D360_Meta struct {
	APIStatus string `json:"api_status"`
	Version   string `json:"version"`
}

type D360_MessageResponse struct {
	Messages []D360_Message `json:"messages"`
	Contacts []D360_Contact `json:"contacts"`
	Meta     D360_Meta      `json:"meta"`
}

type D360_Provider struct {
	Name string `json:"name"`
}

type D360_Document struct {
	Filename string        `json:"filename,omitempty"`
	ID       string        `json:"id,omitempty"`
	Link     string        `json:"link,omitempty"`
	Provider D360_Provider `json:"provider,omitempty"`
}

type D360_Video struct {
	ID       string        `json:"id,omitempty"`
	Link     string        `json:"link,omitempty"`
	Provider D360_Provider `json:"provider,omitempty"`
}

type D360_Image struct {
	ID   string `json:"id,omitempty"`
	Link string `json:"link,omitempty"`
	Text string `json:"caption,omitempty"`
}

type D360_Header struct {
	Type     string        `json:"type"` // "text" | "image" | "video" | "document"
	Text     string        `json:"text,omitempty"`
	Document D360_Document `json:"document,omitempty"`
	Video    D360_Video    `json:"video,omitempty"`
	Image    D360_Image    `json:"image,omitempty"`
}

type D360_Body struct {
	Text string `json:"text,omitempty"`
}

type D360_Footer struct {
	Text string `json:"text,omitempty"`
}

type D360_Reply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type D360_Button struct {
	Type  string     `json:"type"` // "reply" "text"
	Text  string     `json:"text,omitempty"`
	Reply D360_Reply `json:"reply,omitempty"`
}

type D360_Action struct {
	Buttons []D360_Button `json:"buttons"`
}

type D360_Interacrive struct {
	Type   string      `json:"type"` // "button"
	Header D360_Header `json:"header,omitempty"`
	Body   D360_Body   `json:"body,omitempty"`
	Footer D360_Footer `json:"footer,omitempty"`
	Action D360_Action `json:"action,omitempty"`
}

type D360_MessageInteractiveRequest struct {
	RecipientType string           `json:"recipient_type,omitempty"`
	To            string           `json:"to"`
	Type          string           `json:"type"`
	Interactive   D360_Interacrive `json:"interactive,omitempty"`
}

type D360_MessageTemplateRequest struct {
	RecipientType string                 `json:"recipient_type,omitempty"`
	To            string                 `json:"to"`
	Type          string                 `json:"type"`
	Template      D360_TemplateToMessage `json:"template,omitempty"`
}

type D360_TemplateParameter struct {
	Type  string     `json:"type,omitempty"`
	Image D360_Image `json:"image,omitempty"`
}

type D360_TemplateComponent struct {
	Format     string                   `json:"format,omitempty"`
	Text       string                   `json:"text,omitempty"`
	Type       string                   `json:"type"`
	Buttons    []D360_Button            `json:"buttons,omitempty"`
	Parameters []D360_TemplateParameter `json:"parameters,omitempty"`
}

type D360_Template struct {
	Category       string                   `json:"category"`
	Components     []D360_TemplateComponent `json:"components"`
	Language       string                   `json:"language"`
	Name           string                   `json:"name"`
	Namespace      string                   `json:"namespace"`
	RejectedReason string                   `json:"rejected_reason"`
	Status         string                   `json:"status"`
}

type D360_Language struct {
	Policy string `json:"policy,omitempty"`
	Code   string `json:"code,omitempty"`
}

type D360_TemplateToMessage struct {
	Components []D360_TemplateComponent `json:"components"`
	Language   D360_Language            `json:"language"`
	Name       string                   `json:"name"`
	Namespace  string                   `json:"namespace"`
}

type D360_TemplateInteractiveResponse struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Limit         int             `json:"limit"`
	Offset        int             `json:"offset"`
	Sort          []string        `json:"sort"`
	Total         int             `json:"total"`
	WabaTemplates []D360_Template `json:"waba_templates"`
}
