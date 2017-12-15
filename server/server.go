package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	pgx "github.com/jackc/pgx"
)

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
	case "/api/v1/codes":
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
	defer rows.Close()

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

func (s *Server) codes(w http.ResponseWriter, r *http.Request) {
    sql := `UPDATE code SET used = TRUE WHERE id IN (SELECT id FROM code WHERE used = FALSE LIMIT 1) RETURNING code`

    rows, err := s.pool.Query(sql)
    if err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }

    data := []Data{}

    for rows.Next() {
        if values, err := rows.Values(); err == nil {
            data = append(data, Data{
                "idcode":   values[0],
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
