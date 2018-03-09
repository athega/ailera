package server

import (
	"net/http"
	"time"
)

type ListResponse struct {
	Meta   Meta    `json:"meta,omitempty"`
	Data   []Data  `json:"data,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

type Response struct {
	Meta   Meta    `json:"meta,omitempty"`
	Data   Data    `json:"data,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

func writeError(w http.ResponseWriter, r *http.Request, err error, status int, meta Meta) {
	writeJSON(w, errorResponse(err, r.Header.Get("X-Request-ID"), status, meta), status)
}

func errorResponse(err error, id string, status int, meta Meta) Response {
	return Response{
		Meta: meta,
		Errors: []Error{
			{
				ID:     id,
				Status: status,
				Title:  http.StatusText(status),
				Detail: err.Error(),
			},
		},
	}
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

// Error objects provide additional information about problems encountered while performing an operation.
type Error struct {
	ID     string `json:"id,omitempty"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
}
