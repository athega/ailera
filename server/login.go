package server

import (
	"net/http"
	"strings"
)

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.sendLoginEmail(w, r)
		return
	}

	meta := makeMeta(r, s.now())

	id, err := s.storage.ProfileID(r.Context(), r.URL.Query().Get("key"))
	if err != nil {
		writeError(w, r, err, http.StatusUnauthorized, meta)
		return
	}

	token, err := s.signedString(id)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError, meta)
		return
	}

	http.Redirect(w, r, "ailera://Login?token="+token, http.StatusFound)
}

func (s *Server) sendLoginEmail(w http.ResponseWriter, r *http.Request) {
	meta := makeMeta(r, s.now())

	to := r.FormValue("to")

	if !strings.Contains(to, "@") {
		writeError(w, r, errInvalidEmail, http.StatusBadRequest, meta)
		return
	}

	key, err := s.storage.LoginKey(r.Context(), to)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError, meta)
		return
	}

	if err := s.mailer.Send(to,
		"login@ailera.herokuapp.com",
		"Login to ailera",
		`https://ailera.herokuapp.com/login?key=`+key,
		`<a href="https://ailera.herokuapp.com/login?key=`+key+`">Login to ailera</a>`,
	); err != nil {
		writeError(w, r, err, http.StatusInternalServerError, meta)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
