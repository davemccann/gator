package main

import (
	"context"
	"fmt"
)

func command_login(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("invalid argument: command must have 1 argument")
	}

	user, err := s.dbQueries.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("user %s has logged in\n", s.cfg.CurrentUserName)

	return nil
}
