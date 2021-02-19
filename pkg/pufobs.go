package pkg

import (
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

const (
	PUFO        = "\"DAS PODCAST UFO\""
	PUFOFeedURL = "https://rss.acast.com/podcast-ufo"
)

func GetFeed() *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(PUFOFeedURL)
	if err != nil {
		log.Fatal(err)
	}
	return feed
}

func GetEpisodes() []*Episode {
	feed := GetFeed()
	episodes := make([]*Episode, len(feed.Items))
	for i, item := range feed.Items {
		episodes[i] = NewEpisodeFromFeed(item)
	}
	sort.Slice(episodes, func(i, j int) bool {
		return episodes[i].Published.Before(*episodes[j].Published)
	})
	return episodes
}

func GetLatestEpisode() *Episode {
	episodes := GetEpisodes()
	return episodes[len(episodes)-1]
}

func GetEpisode(title string) (*Episode, error) {
	feed := GetFeed()
	for _, item := range feed.Items {
		if strings.HasPrefix(item.Title, title) {
			return NewEpisodeFromFeed(item), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("episode <%s> could not be found", title))
}
