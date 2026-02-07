package main

import (
	"fmt"
	"os"
)

var gateVersion string = "v0.0.1"

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
		defineHelpCommand()
	case "version":
		showGateVersion()
	case "init":
		fmt.Println("gate says Hello")
	case "config":
		fmt.Println("gate says Hello")
	case "run":
		fmt.Println("gate says Hello")
	default:
		return fmt.Errorf("unknown or invalid gate command: %s\nSee `gate help` to see commands", subcommand)
	}

	return nil

}

func defineHelpCommand() {
	fmt.Println("GATE - Automation CLI")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  gate <command>")
	fmt.Println()
	fmt.Println("COMMANDS:")
	fmt.Println("  help            Show this help message")
	fmt.Println("  version         Show gate version")
	fmt.Println()
	fmt.Println("SETUP:")
	fmt.Println("  init            Initialize gate config directory")
	fmt.Println("  config path     Show config directory path")
	fmt.Println("  config list     Show all configuration files")
	fmt.Println()
	fmt.Println("EXECUTION:")
	fmt.Println("  run             Execute steps sequentially")
	fmt.Println()
}

func showGateVersion() {
	fmt.Printf("gate version %s\n", gateVersion)
}
