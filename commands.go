package main

import "fmt"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandFuncs map[string]func(*state, command) error
}

func (cmds *commands) register(name string, fn func(*state, command) error) error {
	if fn == nil {
		return fmt.Errorf("function param to register is nil - name: %s", name)
	}

	if _, exists := cmds.commandFuncs[name]; exists {
		return fmt.Errorf("command already exists in map - name: %s", name)
	}

	cmds.commandFuncs[name] = fn
	return nil
}

func (cmds *commands) run(s *state, cmd command) error {
	fn, exists := cmds.commandFuncs[cmd.name]
	if !exists {
		return fmt.Errorf("could not find function for command - name: %s", cmd.name)
	}

	fn(s, cmd)

	return nil
}

func createCommandsInstance() commands {
	return commands{
		commandFuncs: make(map[string]func(*state, command) error),
	}
}
