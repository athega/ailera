package mail

import "github.com/athega/ailera/ailera"

type loggingMailer struct {
	logger ailera.Logger
}

func NewLoggingMailer(logger ailera.Logger) ailera.Mailer {
	return &loggingMailer{logger: logger}
}

func (lm *loggingMailer) Send(to, from, subject, text, html string) error {
	lm.logger.Printf("Logged mail to=%q from=%q subject=%q text=%q html=%q\n", to, from, subject, text, html)

	return nil
}
