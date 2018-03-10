package mock

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/athega/flockflow-server/flockflow"
)

type Service struct {
	logger flockflow.Logger
}

func NewService(logger flockflow.Logger) *Service {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}

	return &Service{logger: logger}
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
	s.logger.Printf("mail to=%q from=%q subject=%q:\n%s", to, from, subject, text)
	return nil
}
