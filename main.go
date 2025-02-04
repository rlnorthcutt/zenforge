package main

import (
	"os"
)

// Func main should be as small as possible and do as little as possible by convention
func main() {
	// Check if a command is provided, immediately exit if not.
	if len(os.Args) < 2 {
		logger.Warn("Expected a command - type `repose help` to get options.")
		os.Exit(0)
	}

	// Check for flags and set the rootPath and configPath
	command.parseFlags()

	// Get the command name
	commandName := command.Args[0]

	// Dispatch the command
	dispatchCommand(commandName)
}

// **********  Private Main methods  **********

// dispatchCommand will take the command name and dispatch it to the correct function
func dispatchCommand(commandName string) {

	// Load config for specific commands
	switch commandName {
	case "new", "build", "preview":
		var err error
		config, err = config.Load()
		if err != nil {
			logger.Warn("No config file found. You need to run `repose init` first.")
			os.Exit(0)
		}

		// Set rootPath and configPath for command
		// @TODO This is a workaround becuase when we call this in the parse flags
		// the config is not loaded yet... but we can't load it until now.
		buildCommand.SetRootPath(buildCommand.rootPath)
	}

	// Dispatch the command
	switch commandName {
	case "init":
		command.Init()
	case "new":
		command.New(config)
	case "demo":
		command.Demo()
	case "build":
		command.Build(config)
	case "preview":
		command.Preview(config)
	case "update":
		command.Update()
	case "help":
		command.Help()
	default:
		logger.Error("Unknown command: %s\n", os.Args[1])
	}
}
