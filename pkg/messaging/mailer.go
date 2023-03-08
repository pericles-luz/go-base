package messaging

import (
	"errors"
	"log"
	"net/smtp"

	"github.com/pericles-luz/go-base/pkg/utils"
)

type Mailer struct {
	to          string
	subject     string
	body        string
	config      map[string]string
	mimeHeaders string
}

func NewMailer() *Mailer {
	return &Mailer{}
}

func (m *Mailer) SetTo(to string) {
	m.to = to
}

func (m *Mailer) SetSubject(subject string) {
	m.subject = subject
}

func (m *Mailer) SetBody(body string) {
	m.body = body
}

func (m *Mailer) SetConfig(config map[string]string) {
	m.config = config
}

func (m *Mailer) SetMimeHeaders(mimeHeaders string) {
	m.mimeHeaders = mimeHeaders
}

func (m *Mailer) GetTo() string {
	return m.to
}

func (m *Mailer) GetSubject() string {
	return m.subject
}

func (m *Mailer) GetBody() string {
	return m.body
}

func (m *Mailer) GetMimeHeaders() string {
	return m.mimeHeaders
}

func (m *Mailer) Send() error {
	err := m.validate()
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", m.getConfig()["username"], m.getConfig()["password"], m.getConfig()["host"])
	if m.mimeHeaders == "" {
		m.mimeHeaders = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	}
	msg := []byte("Subject: " + m.subject + "\n" + m.mimeHeaders + "\n\n" + m.body)
	log.Println("Sending email to: " + m.to)
	err = smtp.SendMail(m.getConfig()["host"]+":"+m.getConfig()["port"], auth, m.getConfig()["username"], []string{m.to}, msg)
	return err
}

func (m *Mailer) validate() error {
	if !utils.ValidateEmail(m.to) {
		return errors.New("to is required")
	}
	if m.subject == "" {
		return errors.New("subject is required")
	}
	if m.body == "" {
		return errors.New("body is required")
	}
	if m.getConfig() == nil {
		return errors.New("config is required")
	}
	if m.getConfig()["host"] == "" {
		return errors.New("host is required")
	}
	if m.getConfig()["port"] == "" {
		return errors.New("port is required")
	}
	if m.getConfig()["username"] == "" {
		return errors.New("username is required")
	}
	if m.getConfig()["password"] == "" {
		return errors.New("password is required")
	}
	return nil
}

func (m *Mailer) getConfig() map[string]string {
	return m.config
}
