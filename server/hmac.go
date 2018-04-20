package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
)

func (s *Server) toEmailFromRequest(r *http.Request) (string, error) {
	to := r.FormValue("to")

	if !strings.Contains(to, "@") {
		return "", errInvalidEmail
	}

	macString := strings.TrimPrefix(r.Header.Get("Authorization"), "Ailera ")

	mac1, err := decodeHexString(macString)
	if err != nil {
		return "", err
	}

	if !checkMAC([]byte(to), mac1, s.loginKey) {
		return "", errInvalidMAC
	}

	return to, nil
}

func checkMAC(message, mac1, key []byte) bool {
	hash := hmac.New(sha256.New, key)
	hash.Write(message)

	return hmac.Equal(mac1, hash.Sum(nil))
}

func decodeHexString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
