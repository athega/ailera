package flockflow

import "context"

type Service interface {
	UserID(ctx context.Context, key string) (string, error)
	Profile(ctx context.Context, subject string) (*Profile, error)
	SendEmail(to, from, subject, text, html string) error
}

type Logger interface {
	// Printf writes a formated message to the log.
	Printf(format string, v ...interface{})
}
