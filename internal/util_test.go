package internal

import "testing"

func TestFileExists(t *testing.T) {
	if FileExists("doesnotexist.txt") {
		t.Error("doesnotexist.txt should not exist")
	}
	if !FileExists("../README.md") {
		t.Error("README.md should exist")
	}
}
