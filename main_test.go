package main

import (
	"log"
	"os"
	"testing"

	env "github.com/TV4/env"
)

func TestSetup(t *testing.T) {
	logger := log.New(os.Stdout, "", 0)

	hs := setup(logger, env.MapClient(env.Map{}))

	if got, want := hs.Addr, ":3000"; got != want {
		t.Fatalf("hs.Addr = %q, want %q", got, want)
	}
}
