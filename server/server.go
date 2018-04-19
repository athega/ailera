package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"

	"github.com/athega/ailera/ailera"
)

var (
	errHTTPSRequired = errors.New("HTTPS required")
	errInvalidJWT    = errors.New("invalid JWT")
	errInvalidEmail  = errors.New("invalid email")
)

func New(logger ailera.Logger, storage ailera.Store, mailer ailera.Mailer, secretKey []byte) *Server {
	if logger == nil {
		logger = log.New(ioutil.Discard, "", 0)
	}

	s := &Server{
		logger:    logger,
		storage:   storage,
		mailer:    mailer,
		secretKey: secretKey,
	}

	s.registerHandlers()

	return s
}

type Server struct {
	logger    ailera.Logger
	storage   ailera.Store
	mailer    ailera.Mailer
	secretKey []byte
	timeNow   func() time.Time
	mux       *http.ServeMux
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

	s.mux.ServeHTTP(w, r)
}

func (s *Server) registerHandlers() {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}

	s.mux.Handle("/", http.FileServer(http.Dir("docs")))
	s.mux.Handle("/login", http.HandlerFunc(s.login))
	s.mux.Handle("/profile", http.HandlerFunc(s.profile))
	s.mux.Handle("/__status", http.HandlerFunc(s.status))
}

func (s *Server) claimsFromRequest(r *http.Request) (*jwt.StandardClaims, error) {
	var claims jwt.StandardClaims

	token, err := request.ParseFromRequest(r,
		request.OAuth2Extractor, s.keyFunc,
		request.WithClaims(&claims),
		request.WithParser(&jwt.Parser{
			ValidMethods: []string{jwt.SigningMethodHS256.Name},
		}),
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errInvalidJWT
	}

	return &claims, nil
}

func (s *Server) keyFunc(token *jwt.Token) (interface{}, error) {
	return s.secretKey, nil
}

func (s *Server) signedString(sub string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:   "ailera",
		IssuedAt: jwt.TimeFunc().Unix(),
		Subject:  sub,
	}).SignedString([]byte(s.secretKey))
}

func (s *Server) now() time.Time {
	if s.timeNow == nil {
		s.timeNow = time.Now
	}

	return s.timeNow()
}

func (s *Server) log(format string, v ...interface{}) {
	s.logger.Printf(format+"\n", v...)
}

func writeJSON(w http.ResponseWriter, v interface{}, statuses ...int) {
	w.Header().Set("Server", "ailera")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Vary", "Accept-Encoding")

	if len(statuses) > 0 {
		w.WriteHeader(statuses[0])
	}

	enc := json.NewEncoder(w)

	enc.SetIndent("", "  ")
	enc.Encode(v)
}
