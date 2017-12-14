package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var addr = ":3000"

	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	s := NewServer(log.New(os.Stdout, "", 0))

	s.log("Listening on http://localhost%s", addr)

	http.ListenAndServe(addr, s)
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

type Server struct {
	logger *log.Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proto := r.Header.Get("X-Forwarded-Proto")

	if proto == "" {
		proto = "http"
	}

	s.log("%s %s://%s%s", r.Method, proto, r.Host, r.URL.Path)

	w.Write([]byte("FlockFlow"))
}

func (s *Server) log(format string, v ...interface{}) {
	s.logger.Printf(format+"\n", v...)
}
