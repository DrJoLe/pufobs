package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestNewCurrentCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCurrentCmd()
	cmd.SetOut(b)

	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	current := string(out)
	if current == "" {
		t.Fatal("expected an episode name got an empty string")
	}
}
