package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	expected := `pufobs is a small tool to list and download "DAS PODCAST UFO" episodes.

Usage:
  pufobs [flags]

Flags:
  -h, --help   help for pufobs
`

	b := bytes.NewBufferString("")
	cmd := NewRootCmd()
	cmd.SetOut(b)
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}

func TestExecute(t *testing.T) {
	Execute()
}
