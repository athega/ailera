package main

import (
	"log"
	"net/http"
	"os"

	pgx "github.com/jackc/pgx"

	"github.com/athega/flockflow-server/server"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	addr, s := setup(logger)

	logger.Printf("Listening on http://localhost%s\n", addr)

	http.ListenAndServe(addr, s)
}

func setup(logger *log.Logger) (string, *server.Server) {
	var addr = ":3000"

	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	connConfig, err := pgx.ParseConnectionString(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to create connection config: %v\n", err)
	}

	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: 5,
	})
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	return addr, server.NewServer(logger, pool)
}
