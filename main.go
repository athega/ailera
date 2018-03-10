package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	env "github.com/TV4/env"
	graceful "github.com/TV4/graceful"
	sendgrid "github.com/sendgrid/sendgrid-go"

	"github.com/athega/flockflow-server/server"
	"github.com/athega/flockflow-server/service"
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

func setup(logger *log.Logger, e env.Client) *http.Server {
	var (
		port      = e.String("PORT", defaultPort)
		baseURL   = e.URL("BASE_URL", &url.URL{Scheme: "http", Host: "0.0.0.0:" + port})
		service   = service.New(sendgrid.NewSendClient(e.String("SENDGRID_API_KEY", "")), baseURL)
		secretKey = e.Bytes("SECRET_KEY", defaultSecretKey)
	)

	return &http.Server{
		Addr:              ":" + port,
		Handler:           server.New(logger, service, secretKey),
		ReadTimeout:       e.Duration("READ_TIMEOUT", defaultReadTimeout),
		ReadHeaderTimeout: e.Duration("READ_HEADER_TIMEOUT", defaultReadHeaderTimeout),
	}
}
