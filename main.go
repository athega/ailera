package main

import (
	"log"
	"net/http"
	"os"
	"time"

	env "github.com/TV4/env"
	graceful "github.com/TV4/graceful"

	"github.com/athega/ailera/ailera"
	"github.com/athega/ailera/mail"
	"github.com/athega/ailera/server"
	"github.com/athega/ailera/storage"
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

func setup(logger ailera.Logger, e env.Client) *http.Server {
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

	store, err := storage.ConnectAndSetupSchema(e.String("DATABASE_URL",
		"postgres://localhost/ailera_dev?sslmode=disable",
	))
	if err != nil {
		logger.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	return &http.Server{
		Addr:              ":" + port,
		Handler:           server.New(logger, store, mailer, secretKey),
		ReadTimeout:       e.Duration("READ_TIMEOUT", defaultReadTimeout),
		ReadHeaderTimeout: e.Duration("READ_HEADER_TIMEOUT", defaultReadHeaderTimeout),
	}
}
