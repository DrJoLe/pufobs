package pkg

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"time"
)

type Episode struct {
	Id          string
	Title       string
	Description string
	URL         string
	Published   *time.Time
}

func NewEpisodeFromFeed(item *gofeed.Item) *Episode {
	return &Episode{
		Title:       item.Title,
		Description: item.Description,
		URL:         item.Enclosures[0].URL,
		Published:   item.PublishedParsed,
	}
}

func (e Episode) String() string {
	return fmt.Sprintf("%s - %s", e.Title, e.Published.Format(time.RFC3339))
}
