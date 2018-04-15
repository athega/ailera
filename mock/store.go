package mock

import (
	"context"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/athega/flockflow-server/flockflow"
)

type Store struct {
	logger flockflow.Logger
}

func NewStorage(logger flockflow.Logger) flockflow.Store {
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
		return "", flockflow.ErrProfileIDNotFound
	}
}

func (s *Store) Profile(ctx context.Context, subject string) (*flockflow.Profile, error) {
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

func (s *Store) UpdateProfile(ctx context.Context, subject string, v url.Values) (*flockflow.Profile, error) {
	return &flockflow.Profile{}, nil
}
