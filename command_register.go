package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davemccann/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func command_register(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("invalid argument: command must have 1 argument")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	}

	_, err := s.dbQueries.GetUser(context.Background(), userParams.Name)
	if err == nil {
		return fmt.Errorf("user already exists - name: %s", userParams.Name)
	}

	user, err := s.dbQueries.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	fmt.Printf("user has been registered:\n - name: %v\n", user)

	s.cfg.SetUser(user.Name)
	fmt.Printf("updated config with user: %s\n", user.Name)

	return nil
}
