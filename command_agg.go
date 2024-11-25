package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/davemccann/gator/internal/database"
	"github.com/davemccann/gator/internal/rss"
	"github.com/google/uuid"
)

func command_agg(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("invalid number of arguments: agg command requires 1 'time_between_req' argument")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(timeBetweenRequests))
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Printf("Scraping feed...\n\n")
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.dbQueries.GetNextFeed(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Scraping %s...\n", nextFeed.Name)

	err = s.dbQueries.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	feed, feedErr := rss.FetchFeed(context.Background(), nextFeed.Url)
	if feedErr != nil {
		return feedErr
	}

	for _, item := range feed.Channel.Item {

		timePublished, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("unable to determine date published from feed item: \n- name:%s \n- date:%s\n", item.Title, item.PubDate)
			continue
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: timePublished,
			FeedID:      nextFeed.ID,
		}

		_, createErr := s.dbQueries.CreatePost(context.Background(), params)
		if createErr != nil {
			if strings.Contains(createErr.Error(), "duplicate key") {
				continue
			}
			return createErr
		}
	}

	fmt.Printf("Finished scraping %s\n", nextFeed.Name)

	return nil
}
