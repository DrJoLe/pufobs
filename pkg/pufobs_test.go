package pkg

import (
	"testing"
)

var textExpectedTemplate = "expected \"%s\" got \"%s\""

func TestGetFeed(t *testing.T) {
	feed := GetFeed()
	if feed == nil {
		t.Fatalf(textExpectedTemplate, "feed", "nil")
	}
	if feed.Title != "DAS PODCAST UFO" {
		t.Fatalf(textExpectedTemplate, "DAS PODCAST UFO", feed.Title)
	}
}

func TestGetEpisodes(t *testing.T) {
	feed := GetFeed()
	episodes := GetEpisodes()
	if episodes == nil {
		t.Fatalf(textExpectedTemplate, "episodes", "nil")
	}
	if len(episodes) != len(feed.Items) {
		t.Fatalf("episodes got %d elements but feed got %d", len(episodes), len(feed.Items))
	}
	for _, episode := range episodes {
		if episode == nil {
			t.Fatalf("one episode is nil")
		}
	}
}
