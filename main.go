package main

import (
	"log"
	"net/http"
	"os"

	pgx "github.com/jackc/pgx"

	"github.com/athega/flockflow-server/server"
)

func main() {
	var addr = ":3000"

	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	logger := log.New(os.Stdout, "", 0)

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

	s := server.NewServer(logger, pool)

	logger.Printf("Listening on http://localhost%s\n", addr)

	http.ListenAndServe(addr, s)
}
