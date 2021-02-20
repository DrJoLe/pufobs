package cmd

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestNewListCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewListCmd()
	cmd.SetOut(b)

	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	trimmed := strings.TrimSpace(string(out))
	for _, line := range strings.Split(trimmed, "\n") {
		if line[0] != 'U' {
			t.Fatalf("expected line to start with \"u\" got %s", string(line[0]))
		}
	}
}
