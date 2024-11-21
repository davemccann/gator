package main

import (
	"context"
	"fmt"
)

func command_reset(s *state, _ command) error {
	err := s.dbQueries.Reset(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("successfully removed users from the database")

	return nil
}
