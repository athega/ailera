package server

import "net/http"

func (s *Server) profile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.updateProfile(w, r)
		return
	}

	meta := makeMeta(r, s.now())

	claims, err := s.claimsFromRequest(r)
	if err != nil {
		writeError(w, r, err, http.StatusUnauthorized, meta)
		return
	}

	meta["claims"] = claims

	profile, err := s.storage.Profile(r.Context(), claims.Subject)
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

func (s *Server) updateProfile(w http.ResponseWriter, r *http.Request) {
	meta := makeMeta(r, s.now())

	claims, err := s.claimsFromRequest(r)
	if err != nil {
		writeError(w, r, err, http.StatusUnauthorized, meta)
		return
	}

	meta["claims"] = claims

	if err := r.ParseForm(); err != nil {
		writeError(w, r, err, http.StatusBadRequest, meta)
		return
	}

	profile, err := s.storage.UpdateProfile(r.Context(), claims.Subject, r.Form)
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
