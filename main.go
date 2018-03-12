package main

import (
	"log"
	"net/http"
	"os"
	"time"

	env "github.com/TV4/env"
	graceful "github.com/TV4/graceful"

	"github.com/athega/flockflow-server/flockflow"
	"github.com/athega/flockflow-server/mail"
	"github.com/athega/flockflow-server/server"
	"github.com/athega/flockflow-server/storage"
)

var (
	defaultPort              = "3000"
	defaultSecretKey         = []byte("secret")
	defaultReadTimeout       = 20 * time.Second
	defaultReadHeaderTimeout = 10 * time.Second
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	graceful.LogListenAndServe(setup(logger, env.DefaultClient), logger)
}

func setup(logger flockflow.Logger, e env.Client) *http.Server {
	var (
		secretKey = e.Bytes("SECRET_KEY", defaultSecretKey)
		port      = e.String("PORT", defaultPort)
		mailer    = mail.NewLoggingMailer(logger)
	)

	if e.Bool("SEND_EMAIL", false) {
		if key := e.String("SENDGRID_API_KEY", ""); key != "" {
			mailer = mail.NewEmailMailer(key)
		}
	}

	return &http.Server{
		Addr:              ":" + port,
		Handler:           server.New(logger, storage.New(), mailer, secretKey),
		ReadTimeout:       e.Duration("READ_TIMEOUT", defaultReadTimeout),
		ReadHeaderTimeout: e.Duration("READ_HEADER_TIMEOUT", defaultReadHeaderTimeout),
	}
}
