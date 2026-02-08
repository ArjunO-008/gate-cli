package main

import (
	"fmt"
	"gate/cmd"
	"os"
)

func main() {

	if err := runCommand(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

}
func runCommand() error {

	if len(os.Args) < 2 {
		return fmt.Errorf("use: gate <command>")
	}

	return subcommandHandler(os.Args[1])
}

func subcommandHandler(subcommand string) error {

	switch subcommand {
	case "help":
		cmd.DefineHelpCommand()
	case "version":
		cmd.ShowGateVersion()
	case "init":
		cmd.InitCommandHandler()
	case "config":
		fmt.Println("gate says Hello")
	case "run":
		fmt.Println("gate says Hello")
	default:
		return fmt.Errorf("unknown or invalid gate command: %s\nSee `gate help` to see commands", subcommand)
	}

	return nil

}
