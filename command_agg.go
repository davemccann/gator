package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davemccann/blog-aggregator/internal/rss"
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

	err = s.dbQueries.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	feed, feedErr := rss.FetchFeed(context.Background(), nextFeed.Url)
	if feedErr != nil {
		return feedErr
	}

	fmt.Printf("Displaying feed headings for: %s\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Printf("    * %s\n", item.Title)
	}

	return nil
}
