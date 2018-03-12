package flockflow

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrProfileNotFound = errors.New("profile not found")
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type Mailer interface {
	Send(to, from, subject, text, html string) error
}

type Storage interface {
	UserID(ctx context.Context, key string) (string, error)
	Profile(ctx context.Context, subject string) (*Profile, error)
}

type Profile struct {
	ID    string
	Name  string
	Email string
	Link  string
	Phone string
}
