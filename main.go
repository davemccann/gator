package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/davemccann/blog-aggregator/internal/config"
	"github.com/davemccann/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func registerCommands(cmds *commands) error {
	commandMap := map[string]func(*state, command) error{
		"login":     command_login,
		"register":  command_register,
		"reset":     command_reset,
		"users":     command_listusers,
		"agg":       command_agg,
		"addfeed":   authenticateUser(command_addfeed),
		"feeds":     command_listfeeds,
		"follow":    authenticateUser(command_follow),
		"following": authenticateUser(command_following),
		"unfollow":  authenticateUser(command_unfollow),
		"browse":    authenticateUser(command_browse),
	}

	if err := cmds.registerCommands(&commandMap); err != nil {
		return err
	}

	return nil
}

func processCLIArguments(appState *state, cmds *commands) error {
	if len(os.Args) < 2 {
		log.Fatalf("invalid number of arguments")
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	cmd := command{
		name:      commandName,
		arguments: args,
	}

	return cmds.run(appState, cmd)
}

func main() {

	if err := config.EnsureConfigExists(); err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	appState := state{
		dbQueries: dbQueries,
		cfg:       &cfg,
	}

	commands := createCommandsInstance()
	if err := registerCommands(&commands); err != nil {
		log.Fatal(err)
	}

	if err := processCLIArguments(&appState, &commands); err != nil {
		log.Fatal(err)
	}
}
