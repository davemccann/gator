package main

import (
	"context"
	"fmt"
)

func command_following(s *state, _ command) error {
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("no user is currently logged in")
	}

	user, err := s.dbQueries.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
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
