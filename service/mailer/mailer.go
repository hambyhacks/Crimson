package service

import (
	"context"

	klog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hambyhacks/CrimsonIMS/app/models"
)

type MailerService interface {
	SendMail(ctx context.Context, email models.Mailer) (string, error)
}

type MailerServ struct {
	mailrepo MailRepository
	logger   klog.Logger
}

func NewMailServ(mailrepo MailRepository, logger klog.Logger) MailerService {
	return &MailerServ{
		mailrepo: mailrepo,
		logger:   logger,
	}
}

// SendMail implements MailerService
func (m *MailerServ) SendMail(ctx context.Context, email models.Mailer) (string, error) {
	logger := klog.With(m.logger, "method", "send mail")
	msg := "email sent successfully"
	emailDetails := &models.Mailer{
		Recipient:    email.Recipient,
		TemplateFile: email.TemplateFile,
		Data:         email.Data,
	}

	_, err := m.mailrepo.SendMail(ctx, *emailDetails)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return "unable to process request", err
	}
	return msg, nil
}
