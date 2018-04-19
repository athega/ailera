package mock

import (
	"context"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/athega/ailera/ailera"
)

type Store struct {
	logger ailera.Logger
}

func NewStorage(logger ailera.Logger) ailera.Store {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}

	return &Store{logger: logger}
}

func (s *Store) LoginKey(ctx context.Context, email string) (string, error) {
	return "1234", nil
}

func (s *Store) ProfileID(ctx context.Context, key string) (string, error) {
	switch key {
	case "1234":
		return "5678", nil
	default:
		return "", ailera.ErrProfileIDNotFound
	}
}

func (s *Store) Profile(ctx context.Context, subject string) (*ailera.Profile, error) {
	switch subject {
	case "5678":
		return &ailera.Profile{
			ID:    "5678",
			Name:  "Foo Bar",
			Email: "foo.bar@example.com",
			Link:  "http://example.com/",
			Phone: "012345678",
		}, nil
	default:
		return nil, ailera.ErrProfileNotFound
	}
}

func (s *Store) UpdateProfile(ctx context.Context, subject string, v url.Values) (*ailera.Profile, error) {
	return &ailera.Profile{}, nil
}
