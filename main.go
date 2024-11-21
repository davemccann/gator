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
	commandMappings := []struct {
		commandName string
		callbackFn  func(*state, command) error
	}{
		{
			commandName: "login",
			callbackFn:  command_login,
		},
		{
			commandName: "register",
			callbackFn:  command_register,
		},
		{
			commandName: "reset",
			callbackFn:  command_reset,
		},
		{
			commandName: "users",
			callbackFn:  command_listusers,
		},
		{
			commandName: "agg",
			callbackFn:  command_agg,
		},
		{
			commandName: "addfeed",
			callbackFn:  command_addfeed,
		},
	}

	for _, mapping := range commandMappings {
		if err := cmds.register(mapping.commandName, mapping.callbackFn); err != nil {
			return err
		}
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
