package main

import (
	"context"
	"fmt"

	"github.com/davemccann/blog-aggregator/internal/database"
)

func command_listfeeds(s *state, _ command) error {
	feeds, err := s.dbQueries.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.dbQueries.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Printf("Could not find user in database with ID: %s", feed.UserID)
			continue
		}

		printDatabaseFeed(&feed, user.Name)
	}

	return nil
}

func printDatabaseFeed(feed *database.Feed, createdBy string) {
	outputFormat := `
--------------------------
* Name:            %s
* URL:             %s
* Created By:      %s
--------------------------
`
	fmt.Printf(outputFormat, feed.Name, feed.Url, createdBy)
}
