package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/davemccann/blog-aggregator/internal/config"
	"github.com/davemccann/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func registerCommands(cmds *commands) error {

	if err := cmds.register("login", command_login); err != nil {
		return err
	}

	if err := cmds.register("register", command_register); err != nil {
		return err
	}

	if err := cmds.register("reset", command_reset); err != nil {
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

	fmt.Printf("Config Output: - CurrentUserName: %s - DbURL %s\n", appState.cfg.CurrentUserName, appState.cfg.DbURL)
}
