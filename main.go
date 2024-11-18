package main

import (
	"fmt"
	"log"

	"github.com/davemccann/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Config Output: - CurrentUserName: %s - DbURL %s\n", cfg.CurrentUserName, cfg.DbURL)
	cfg.SetUser("John")
	fmt.Printf("Config Output: - CurrentUserName: %s - DbURL %s\n", cfg.CurrentUserName, cfg.DbURL)
}
