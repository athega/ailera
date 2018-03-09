package server

import "net/http"

func (s *Server) profile(w http.ResponseWriter, r *http.Request) {
	s.log("Authorization: %#v", r.Header.Get("Authorization"))

	meta := makeMeta(r, s.now())

	writeJSON(w, Response{
		Meta: meta,
		Data: Data{
			"foo": "bar",
		},
	})
}
