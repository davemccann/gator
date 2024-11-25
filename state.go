package main

import (
	"github.com/davemccann/gator/internal/config"
	"github.com/davemccann/gator/internal/database"
)

type state struct {
	dbQueries *database.Queries
	cfg       *config.Config
}
