package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx"
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

	s := NewServer(logger, pool)

	s.log("Listening on http://localhost%s", addr)

	http.ListenAndServe(addr, s)
}

func NewServer(logger *log.Logger, pool *pgx.ConnPool) *Server {
	return &Server{
		logger: logger,
		pool:   pool,
	}
}

type Server struct {
	logger *log.Logger
	pool   *pgx.ConnPool
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proto := r.Header.Get("X-Forwarded-Proto")

	if proto == "" {
		proto = "http"
	}

	s.log("%s %s://%s%s", r.Method, proto, r.Host, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		s.get(w, r)
	}
}

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/v1/db":
		s.db(w, r)
	default:
		w.Write([]byte("FlockFlow"))
	}
}

func (s *Server) db(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT firstname, lastname, image, pers_number, email FROM userdata`

	rows, err := s.pool.Query(sql)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := []Data{}

	for rows.Next() {
		if values, err := rows.Values(); err == nil {
			data = append(data, Data{
				"firstname":   values[0],
				"lastname":    values[1],
				"image":       values[2],
				"pers_number": values[3],
				"email":       values[4],
			})
		}
	}

	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(ListResponse{
		Meta: makeMeta(r, time.Now()),
		Data: data,
	})
}

func (s *Server) log(format string, v ...interface{}) {
	s.logger.Printf(format+"\n", v...)
}

type ListResponse struct {
	Meta Meta   `json:"meta,omitempty"`
	Data []Data `json:"data,omitempty"`
}

type Response struct {
	Meta Meta `json:"meta,omitempty"`
	Data Data `json:"data,omitempty"`
}

func makeMeta(r *http.Request, now time.Time, options ...func(Meta)) Meta {
	meta := Meta{
		"request_id":    r.Header.Get("X-Request-ID"),
		"forwarded_for": r.Header.Get("X-Forwarded-For"),
		"timestamp":     now.UTC(),
		"endpoint":      r.URL.Path,
	}

	for _, f := range options {
		f(meta)
	}

	return meta
}

// Meta is a meta object that contains non-standard meta-information.
type Meta map[string]interface{}

// Data is the primary data of the Response
type Data map[string]interface{}
