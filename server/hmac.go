package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func checkMAC(message, mac1, key []byte) bool {
	hash := hmac.New(sha256.New, key)
	hash.Write(message)

	return hmac.Equal(mac1, hash.Sum(nil))
}

func decodeURLEncodedString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
