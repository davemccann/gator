package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davemccann/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func command_addfeed(s *state, cmd command, user *database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("invalid number of arguments: addfeed requires 'feed name' and 'url'")
	}

	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}

	feedEntry, err := s.dbQueries.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	fmt.Println("Successfully created feed:")
	printFeed(&feedEntry)

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedEntry.ID,
	}

	feedsFollow, err := s.dbQueries.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println("Successfully created feeds follow:")
	printFeedsFollow(&feedsFollow)

	return nil
}

func printFeed(feed *database.Feed) {
	outputFormat := `
	* ID:			%s
	* Created:		%v
	* Updated:		%v
	* Name:			%s
	* URL:			%s
	* UserID:		%s

`
	fmt.Printf(outputFormat, feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, feed.UserID)
}
