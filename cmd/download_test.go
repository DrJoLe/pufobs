package cmd

import (
	"bytes"
	"github.com/drjole/pufobs/internal"
	"io"
	"os"
	"testing"
)

func TestNewDownloadCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewDownloadCmd()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"UFO001 Prolog", "pufo.mp3"})

	if internal.FileExists("pufo.mp3") {
		_ = os.Remove("pufo.mp3")
	}

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

	if !internal.FileExists("pufo.mp3") {
		t.Fatal("download failed")
	}

	stat, err := os.Stat("pufo.mp3")
	if err != nil {
		t.Fatal(err)
	}

	if stat.Size() <= 50000000 {
		t.Fatalf("file is not at least 50000000 bytes long but %d", stat.Size())
	}

	_ = os.Remove("pufo.mp3")
}
