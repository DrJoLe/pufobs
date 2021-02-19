package pkg

import (
	"github.com/mmcdole/gofeed"
	"testing"
	"time"
)

func TestNewEpisodeFromFeed(t *testing.T) {
	feedItem := &gofeed.Item{
		Title:       "Title",
		Description: "Description",
		Enclosures: []*gofeed.Enclosure{
			{
				URL: "URL",
			},
		},
		PublishedParsed: &[]time.Time{time.Unix(0, 0)}[0],
	}
	episode := NewEpisodeFromFeed(feedItem)
	if episode.Title != feedItem.Title {
		t.Fatalf(textExpectedTemplate, feedItem.Title, episode.Title)
	}
	if episode.Description != feedItem.Description {
		t.Fatalf(textExpectedTemplate, feedItem.Title, episode.Title)
	}
	if episode.URL != feedItem.Enclosures[0].URL {
		t.Fatalf(textExpectedTemplate, feedItem.Title, episode.Title)
	}
	if episode.Published != feedItem.PublishedParsed {
		t.Fatalf(textExpectedTemplate, feedItem.Title, episode.Title)
	}
}

func TestEpisode_String(t *testing.T) {
	episode := Episode{
		"ID",
		"Title",
		"Description",
		"https://example.com",
		&[]time.Time{time.Unix(0, 0)}[0],
	}
	expected := "Title - 1970-01-01T01:00:00+01:00"
	if episode.String() != expected {
		t.Fatalf(textExpectedTemplate, expected, episode.String())
	}
}
