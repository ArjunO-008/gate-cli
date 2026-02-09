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

	args := os.Args[2:]
	return subcommandHandler(os.Args[1], args)
}

func subcommandHandler(subcommand string, arguments []string) error {

	switch subcommand {
	case "help":
		cmd.DefineHelpCommand()
	case "version":
		cmd.ShowGateVersion()
	case "init":
		if err := cmd.InitCommandHandler(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "config":
		if err := cmd.ConfigCommandHandler(arguments); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "run":
		fmt.Println("gate says Hello")
	default:
		return fmt.Errorf("unknown or invalid gate command: %s\nSee `gate help` to see commands", subcommand)
	}

	return nil

}
