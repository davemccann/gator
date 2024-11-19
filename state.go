package main

import (
	"github.com/davemccann/blog-aggregator/internal/config"
	"github.com/davemccann/blog-aggregator/internal/database"
)

type state struct {
	dbQueries *database.Queries
	cfg       *config.Config
}
