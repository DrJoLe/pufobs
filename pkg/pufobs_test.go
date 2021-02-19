package pkg

import (
	"testing"
)

var textExpectedTemplate = "expected \"%s\" got \"%s\""

func TestGetFeed(t *testing.T) {
	feed := GetFeed()
	if feed == nil {
		t.Fatalf("feed is nil")
	}
	if feed.Title != "DAS PODCAST UFO" {
		t.Fatalf(textExpectedTemplate, "DAS PODCAST UFO", feed.Title)
	}
}
