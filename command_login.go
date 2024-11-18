package main

import "fmt"

func command_login(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("invalid argument: command must have 1 argument")
	}

	if err := s.cfg.SetUser(cmd.arguments[0]); err != nil {
		return err
	}

	fmt.Printf("username has been set to %s\n", s.cfg.CurrentUserName)

	return nil
}
