package server

import (
	"bytes"
	"log"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	if got := New(nil, "testsecret"); got == nil {
		t.Fatal("New returned nil")
	}
}

func TestServeHTTP(t *testing.T) {
	var buf bytes.Buffer

	logger := log.New(&buf, "", 0)

	s := New(logger, "testsecret")

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
