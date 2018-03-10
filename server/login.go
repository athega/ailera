package server

import "net/http"

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//s.sendLoginEmail(w, r)
		return
	}

	meta := makeMeta(r, s.now())

	key := r.URL.Query().Get("key")

	id, err := s.service.UserID(r.Context(), key)
	if err != nil {
		writeError(w, r, err, http.StatusUnauthorized, meta)
		return
	}

	token, err := s.signedString(id)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError, meta)
		return
	}

	http.Redirect(w, r, "flockflow://Login?token="+token, http.StatusFound)
}

func (s *Server) sendLoginEmail(w http.ResponseWriter, r *http.Request) {
	meta := makeMeta(r, s.now())

	if err := s.service.SendEmail(
		"peter.hellberg@athega.se",
		"login@flockflow.herokuapp.com",
		"Login to FlockFlow",
		`https://flockflow.herokuapp.com/login?key=1234`,
		`<a href="https://flockflow.herokuapp.com/login?key=1234">Login to FlockFlow</a>`,
	); err != nil {
		writeError(w, r, err, http.StatusInternalServerError, meta)
	}

	w.WriteHeader(http.StatusAccepted)
}
