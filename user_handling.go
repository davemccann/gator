package main

import (
	"context"
	"fmt"

	"github.com/davemccann/blog-aggregator/internal/database"
)

func authenticateUser(handler func(s *state, cmd command, user *database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {

		user, err := s.dbQueries.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to fetch active user ensure you have registered a user or logged in - error: %v", err)
		}

		return handler(s, cmd, &user)
	}
}
