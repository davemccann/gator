package main

import (
	"context"
	"fmt"

	"github.com/davemccann/blog-aggregator/internal/database"
)

func command_unfollow(s *state, cmd command, user *database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("invalid arguments: unfollow requires 'feed url'")
	}

	feedURL := cmd.arguments[0]

	deleteFeedFollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    feedURL,
	}

	result, err := s.dbQueries.DeleteFeedFollow(context.Background(), deleteFeedFollowParams)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		fmt.Printf("No feeds unfollowed using url: %s", feedURL)
	}

	return nil
}
