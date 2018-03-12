package server

import (
	"bytes"
	"log"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/athega/flockflow-server/flockflow"
	"github.com/athega/flockflow-server/mock"
)

func TestNew(t *testing.T) {
	if got := New(nil, mock.NewStorage(nil), mock.NewMailer(nil), testSecretKey); got == nil {
		t.Fatal("New returned nil")
	}
}

func TestServeHTTP(t *testing.T) {
	var buf bytes.Buffer

	logger := log.New(&buf, "", 0)

	s := testServer(logger)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://example.org/__status", nil)

	s.ServeHTTP(w, r)

	if got, want := w.Code, 200; got != want {
		t.Fatalf("w.Code = %d, want %d", got, want)
	}

	if got, want := buf.String(), "GET http://example.org/__status\n"; got != want {
		t.Fatalf("buf.String() = %q, want %q", got, want)
	}
}

var testSecretKey = []byte("testsecret")

func testServer(logger flockflow.Logger, options ...func(*Server)) *Server {
	s := New(logger, mock.NewStorage(logger), mock.NewMailer(logger), testSecretKey)

	s.timeNow = func() time.Time {
		return time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
	}

	for _, f := range options {
		f(s)
	}

	return s
}
