package server

import "net/http"

var hardcodedLoginKeys = map[string]string{
	"1234": "5678",
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	id, ok := hardcodedLoginKeys[r.URL.Query().Get("key")]
	if !ok {
		writeError(w, r, errUnknownKey, http.StatusUnauthorized, makeMeta(r, s.now()))
		return
	}

	token, err := s.signedString(id)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError, makeMeta(r, s.now()))
		return
	}

	http.Redirect(w, r, "flockflow://Login?token="+token, http.StatusPermanentRedirect)
}
