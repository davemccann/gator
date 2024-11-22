package main

import (
	"context"
	"fmt"

	"github.com/davemccann/blog-aggregator/internal/database"
)

func command_following(s *state, _ command, user *database.User) error {
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("no user is currently logged in")
	}

	feeds, err := s.dbQueries.GetFeedFollowsByUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("\n Feeds followed by %s:\n", user.Name)
	for _, feedFollowsRow := range feeds {
		fmt.Printf("	* Feed Name: %s\n", feedFollowsRow.FeedName)
	}
	fmt.Println("")

	return nil
}
