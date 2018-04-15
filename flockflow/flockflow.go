package flockflow

import (
	"context"
	"errors"
	"net/url"
	"time"
)

var (
	ErrInvalidLoginKey   = errors.New("invalid login key")
	ErrProfileIDNotFound = errors.New("profile id not found")
	ErrProfileNotFound   = errors.New("profile not found")
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type Mailer interface {
	Send(to, from, subject, text, html string) error
}

type Store interface {
	LoginKey(ctx context.Context, email string) (string, error)
	ProfileID(ctx context.Context, key string) (string, error)
	Profile(ctx context.Context, subject string) (*Profile, error)
	UpdateProfile(ctx context.Context, subject string, v url.Values) error
}

type Profile struct {
	ID    string
	Email string
	Name  string
	Link  string
	Phone string
}

type Login struct {
	Key       string
	Email     string
	Timestamp time.Time
}
