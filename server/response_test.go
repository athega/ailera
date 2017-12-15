package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMakeMeta(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/foo/bar", nil)

	r.Header.Set("X-Request-ID", "test-id-123")
	r.Header.Set("X-Forwarded-For", "0.0.0.0")

	now := time.Date(2017, 12, 15, 10, 0, 0, 0, time.UTC)

	m := makeMeta(r, now, func(m Meta) {
		m["xyz"] = 123.456
	})

	if reqID, ok := m["request_id"].(string); ok {
		if got, want := reqID, "test-id-123"; got != want {
			t.Fatalf(`m["request_id"] = %q, want %q`, got, want)
		}
	} else {
		t.Fatal(`m["request_id"] is not a string`)
	}

	if endpoint, ok := m["endpoint"].(string); ok {
		if got, want := endpoint, "/foo/bar"; got != want {
			t.Fatalf(`m["endpoint"] = %q, want %q`, got, want)
		}
	} else {
		t.Fatal(`m["endpoint"] is not a string`)
	}
}
