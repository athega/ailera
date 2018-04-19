package mock

import (
	"io/ioutil"
	"log"

	"github.com/athega/ailera/ailera"
)

type Mailer struct {
	logger ailera.Logger
}

func NewMailer(logger ailera.Logger) ailera.Mailer {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}

	return &Mailer{logger: logger}
}

func (m *Mailer) Send(to, from, subject, text, html string) error {
	m.logger.Printf("Logged mail to=%q from=%q subject=%q text=%q html=%q\n", to, from, subject, text, html)

	return nil
}
