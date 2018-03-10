package server

import "net/http"

func (s *Server) profile(w http.ResponseWriter, r *http.Request) {
	meta := makeMeta(r, s.now())

	claims, err := s.claimsFromRequest(r)
	if err != nil {
		writeError(w, r, err, http.StatusUnauthorized, meta)
		return
	}

	meta["claims"] = claims

	profile, err := s.service.Profile(r.Context(), claims.Subject)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError, meta)
		return
	}

	writeJSON(w, Response{
		Meta: meta,
		Data: Data{
			"id":    profile.ID,
			"name":  profile.Name,
			"email": profile.Email,
			"link":  profile.Link,
			"phone": profile.Phone,
		},
	})
}
