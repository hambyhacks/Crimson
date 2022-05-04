package service

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"time"

	klog "github.com/go-kit/log"
	"github.com/go-mail/mail/v2"

	"github.com/hambyhacks/CrimsonIMS/app/models"
)

// Error messages
var (
	ErrMailNotSent         = errors.New("unable to send email")
	ErrCannotParseTemplate = errors.New("unable to parse template")
)

// Mail Repository Interface
type MailRepository interface {
	SendMail(ctx context.Context, email models.Mailer) (string, error)
}

type mailRepo struct {
	mailer models.MailDialer
	logger klog.Logger
}

func NewMailRepo(mailer models.MailDialer, logger klog.Logger) (MailRepository, error) {
	return &mailRepo{
		mailer: mailer,
		logger: klog.With(logger, "repo", "mailtrap"),
	}, nil
}

// SendMail implements MailRepository
func (m *mailRepo) SendMail(ctx context.Context, email models.Mailer) (string, error) {
	tmpl, err := template.New("email").ParseFS(models.TemplateFS, "templates/email"+email.TemplateFile)
	if err != nil {
		return "unable to process request", ErrCannotParseTemplate
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", email.Data)
	if err != nil {
		return "unable to process request", ErrCannotParseTemplate
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", email.Data)
	if err != nil {
		return "unable to process request", ErrCannotParseTemplate
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", email.Data)
	if err != nil {
		return "unable to process request", ErrCannotParseTemplate
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", email.Recipient)
	msg.SetHeader("From", m.mailer.Sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	for i := 1; i <= 3; i++ {
		err = m.mailer.Dialer.DialAndSend()
		if nil == err {
			return "", nil
		}
		time.Sleep(500 * time.Millisecond)
	}

	return "unable to process request", ErrMailNotSent
}
