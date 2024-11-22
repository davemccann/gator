package main

import (
	"context"

	"github.com/davemccann/blog-aggregator/internal/database"
)

func authenticateUser(handler func(s *state, cmd command, user *database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {

		user, err := s.dbQueries.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, &user)
	}
}
