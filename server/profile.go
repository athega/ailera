package server

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

type Profile struct {
	Name string
}

var hardcodedProfiles = map[string]Profile{
	"5678": Profile{
		Name: "Foo Bar",
	},
}

func (s *Server) profile(w http.ResponseWriter, r *http.Request) {
	meta := makeMeta(r, s.now())

	s.log("DEBUG: Authorization: %#v", r.Header.Get("Authorization"))

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, s.keyFunc)
	if err != nil {
		writeError(w, r, err, http.StatusUnauthorized, meta)
		return
	}

	if !token.Valid || token.Method != jwt.SigningMethodHS256 {
		writeError(w, r, errInvalidJWT, http.StatusUnauthorized, meta)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		writeError(w, r, errInvalidJWTClaims, http.StatusBadRequest, meta)
		return
	}

	writeJSON(w, Response{
		Meta: meta,
		Data: Data{
			"claims": claims,
		},
	})
}
