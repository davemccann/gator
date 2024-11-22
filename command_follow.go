package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davemccann/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func command_follow(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("invalid arguments: listfeeds requires a URL as an argument")
	}

	feedURL := cmd.arguments[0]

	user, err := s.dbQueries.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.dbQueries.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedsFollow, err := s.dbQueries.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	printFeedsFollow(&feedsFollow)

	return nil
}

func printFeedsFollow(feedsFollow *database.CreateFeedFollowRow) {
	outputFormat := `
	* ID:            %s
	* FeedName:      %s
	* UserName:      %s

`

	fmt.Printf(outputFormat, feedsFollow.ID, feedsFollow.FeedName, feedsFollow.UserName)
}
