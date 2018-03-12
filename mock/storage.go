package mock

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/athega/flockflow-server/flockflow"
)

type Storage struct {
	logger flockflow.Logger
}

func NewStorage(logger flockflow.Logger) flockflow.Storage {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}

	return &Storage{logger: logger}
}

func (s *Storage) UserID(ctx context.Context, key string) (string, error) {
	switch key {
	case "1234":
		return "5678", nil
	default:
		return "", flockflow.ErrUserNotFound
	}
}

func (s *Storage) Profile(ctx context.Context, subject string) (*flockflow.Profile, error) {
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
		return nil, flockflow.ErrProfileNotFound
	}
}
