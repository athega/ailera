package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

var (
	errHTTPSRequired   = errors.New("HTTPS required")
	errUnknownKey      = errors.New("unknown key")
	errUnknownEndpoint = errors.New("unknown endpoint")
)

func New(logger *log.Logger, secretKey string) *Server {
	return &Server{
		logger:    logger,
		secretKey: secretKey,
	}
}

type Server struct {
	logger    *log.Logger
	secretKey string
	timeNow   func() time.Time
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Forwarded-Proto") == "http" {
		writeError(w, r, errHTTPSRequired, http.StatusForbidden, nil)
		return
	}

	proto := r.Header.Get("X-Forwarded-Proto")

	if proto == "" {
		proto = "http"
	}

	s.log("%s %s://%s%s", r.Method, proto, r.Host, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		s.get(w, r)
	case http.MethodPost:
		s.post(w, r)
	}
}

func (s *Server) now() time.Time {
	if s.timeNow == nil {
		s.timeNow = time.Now
	}

	return s.timeNow()
}

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	meta := makeMeta(r, s.now())

	switch r.URL.Path {
	case "/":
		s.index(w, r)
	case "/login":
		s.login(w, r)
	case "/profile":
		s.profile(w, r)
	default:
		writeError(w, r, errUnknownEndpoint, http.StatusBadRequest, meta)
	}
}

func (s *Server) post(w http.ResponseWriter, r *http.Request) {}

func (s *Server) log(format string, v ...interface{}) {
	s.logger.Printf(format+"\n", v...)
}

func writeJSON(w http.ResponseWriter, v interface{}, statuses ...int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Vary", "Accept-Encoding")

	if len(statuses) > 0 {
		w.WriteHeader(statuses[0])
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	enc.Encode(v)
}
