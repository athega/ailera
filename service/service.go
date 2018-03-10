package service

import (
	"context"
	"net/url"

	sendgrid "github.com/sendgrid/sendgrid-go"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/athega/flockflow-server/flockflow"
)

type Service struct {
	sendgrid *sendgrid.Client
	baseURL  *url.URL
}

func New(sc *sendgrid.Client, baseURL *url.URL) *Service {
	return &Service{sendgrid: sc, baseURL: baseURL}
}

func (s *Service) UserID(ctx context.Context, key string) (string, error) {
	switch key {
	case "1234":
		return "5678", nil
	default:
		return "", flockflow.ErrUserNotFound
	}
}

func (s *Service) Profile(ctx context.Context, subject string) (*flockflow.Profile, error) {
	switch subject {
	case "5678":
		return &flockflow.Profile{
			ID:    "5678",
			Name:  "Foo Bar",
			Email: "foo.bar@example.com",
			Link:  "http://example.com/",
			Phone: "012345678",
		}, nil
	default:
		return nil, nil
	}
}

func (s *Service) SendEmail(to, from, subject, text, html string) error {
	email := mail.NewSingleEmail(mail.NewEmail(from, from), subject, mail.NewEmail(to, to), text, html)

	_, err := s.sendgrid.Send(email)

	return err
}
