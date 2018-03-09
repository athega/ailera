package server

import "net/http"

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, Response{
		Meta: makeMeta(r, s.now()),
		Data: Data{},
	})
}
