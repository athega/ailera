package main

import (
	"log"
	"os"
	"testing"
)

func TestSetup(t *testing.T) {
	logger := log.New(os.Stdout, "", 0)

	addr, s := setup(logger)

	if got, want := addr, ":3000"; got != want {
		t.Fatalf("addr = %q, want %q", got, want)
	}

	if s == nil {
		t.Fatal("s is nil")
	}
}
