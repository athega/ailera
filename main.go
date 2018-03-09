package main

import (
	"log"
	"net/http"
	"os"
	"time"

	env "github.com/TV4/env"
	graceful "github.com/TV4/graceful"

	"github.com/athega/flockflow-server/server"
)

const (
	defaultPort              = "3000"
	defaultReadTimeout       = 20 * time.Second
	defaultReadHeaderTimeout = 10 * time.Second
)

var defaultSecretKey = []byte("secret")

func main() {
	logger := log.New(os.Stdout, "", 0)

	graceful.LogListenAndServe(setup(logger, env.DefaultClient), logger)
}

func setup(logger *log.Logger, e env.Client) *http.Server {
	return &http.Server{
		Addr:              ":" + e.String("PORT", defaultPort),
		Handler:           server.New(logger, e.Bytes("SECRET_KEY", defaultSecretKey)),
		ReadTimeout:       e.Duration("READ_TIMEOUT", defaultReadTimeout),
		ReadHeaderTimeout: e.Duration("READ_HEADER_TIMEOUT", defaultReadHeaderTimeout),
	}
}
