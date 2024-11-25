package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davemccann/gator/internal/database"
	"github.com/google/uuid"
)

func command_register(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("invalid argument: command must have 1 argument")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
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

	printUser(&user)

	s.cfg.SetUser(user.Name)
	fmt.Printf("updated config with user: %s\n", user.Name)

	return nil
}

func printUser(user *database.User) {
	outputFormat := `
User has been registered:
* ID:           %s
* Name:         %s
* Created:      %v
* Updated:      %v

`
	fmt.Printf(outputFormat, user.ID, user.Name, user.CreatedAt, user.UpdatedAt)
}
