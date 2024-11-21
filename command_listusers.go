package main

import (
	"context"
	"fmt"
)

func command_listusers(s *state, cmd command) error {
	users, err := s.dbQueries.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {

		current := ""
		if s.cfg.CurrentUserName == user.Name {
			current = "(current)"
		}
		fmt.Printf("* %s %s\n", user.Name, current)
	}

	return nil
}
