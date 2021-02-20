package cmd

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"testing"
)

func TestNewCountCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCountCmd()
	cmd.SetOut(b)

	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	number, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		t.Fatal(err)
	}

	if number <= 0 {
		t.Fatalf("expected %s got %d", "at least 1 episode", number)
	}
}
