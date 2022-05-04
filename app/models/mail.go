package models

import (
	"embed"

	"github.com/go-mail/mail"
)

//go:embed"templates"
var TemplateFS embed.FS

type Mailer struct {
	Recipient    string
	TemplateFile string
	Data         interface{}
}

type MailDialer struct {
	Dialer *mail.Dialer
	Sender string
}
