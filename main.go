package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davemccann/blog-aggregator/internal/config"
)

func registerCommands(cmds *commands) error {
	return cmds.register("login", command_login)
}

func processCLIArguments(appState *state, cmds *commands) error {
	if len(os.Args) < 2 {
		log.Fatalf("invalid number of arguments")
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	if len(args) == 0 {
		log.Fatalf("invalid number of arguments for command: %s", commandName)
	}

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

	appState := state{
		cfg: &cfg,
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
