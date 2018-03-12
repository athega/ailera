package mail

import "github.com/athega/flockflow-server/flockflow"

type loggingMailer struct {
	logger flockflow.Logger
}

func NewLoggingMailer(logger flockflow.Logger) flockflow.Mailer {
	return &loggingMailer{logger: logger}
}

func (lm *loggingMailer) Send(to, from, subject, text, html string) error {
	lm.logger.Printf("Logged mail to=%q from=%q subject=%q text=%q html=%q\n", to, from, subject, text, html)

	return nil
}
