package server

import (
	"bytes"
	"log"
	"net/http/httptest"
	"testing"
)

func TestNewServer(t *testing.T) {
	if got := NewServer(nil, nil); got == nil {
		t.Fatal("NewServer returned nil")
	}
}

func TestServeHTTP(t *testing.T) {
	var buf bytes.Buffer

	logger := log.New(&buf, "", 0)

	s := NewServer(logger, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://example.org/", nil)

	s.ServeHTTP(w, r)

	if got, want := w.Code, 200; got != want {
		t.Fatalf("w.Code = %d, want %d", got, want)
	}

	if got, want := buf.String(), "GET http://example.org/\n"; got != want {
		t.Fatalf("buf.String() = %q, want %q", got, want)
	}
}
