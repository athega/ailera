package server

import (
	"net/http"
	"time"
)

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
