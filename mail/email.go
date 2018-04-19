package mail

import (
	sendgrid "github.com/sendgrid/sendgrid-go"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/athega/ailera/ailera"
)

type emailMailer struct {
	client *sendgrid.Client
}

func NewEmailMailer(key string) ailera.Mailer {
	return &emailMailer{sendgrid.NewSendClient(key)}
}

func (em *emailMailer) Send(to, from, subject, text, html string) error {
	email := mail.NewSingleEmail(mail.NewEmail(from, from), subject, mail.NewEmail(to, to), text, html)

	_, err := em.client.Send(email)

	return err
}
