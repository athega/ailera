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

	// TODO: REPLACE WITH JWT
	token := id

	http.Redirect(w, r, "flockflow://Login?token="+token, http.StatusPermanentRedirect)
}
