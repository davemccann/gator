package main

import (
	"context"
	"fmt"

	"github.com/davemccann/blog-aggregator/internal/rss"
)

const (
	testURL = "https://www.wagslane.dev/index.xml" // TODO(Dave): Temp to test feed will be remove in future
)

func command_agg(_ *state, _ command) error {
	feed, err := rss.FetchFeed(context.Background(), testURL)
	if err != nil {
		return err
	}

	fmt.Printf("Feed description: %v\n", feed)
	return nil
}
