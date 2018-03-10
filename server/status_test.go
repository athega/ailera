package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	s := testServer(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/__status", nil)

	s.status(w, r)

	if got, want := w.Code, http.StatusOK; got != want {
		t.Fatalf("w.Code = %d, want %d", got, want)
	}

	var response struct {
		Data struct {
			Language string `json:"language"`
		} `json:"data"`
	}

	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got, want := response.Data.Language, "go"; got != want {
		t.Fatalf("response.Data.Language = %q, want %q", got, want)
	}
}
