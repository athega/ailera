package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	var msg, key string

	flag.StringVar(&msg, "msg", "", "Message to generate HMAC for")
	flag.StringVar(&key, "key", "", "Key to use in the HMAC")

	flag.Parse()

	if msg == "" || key == "" {
		return
	}

	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(msg))

	fmt.Println("HMAC for", msg, "using key", key, "=",
		base64.StdEncoding.EncodeToString(hash.Sum(nil)),
	)
}
