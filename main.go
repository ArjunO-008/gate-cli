package main

import (
	"fmt"
	"os"
	"path/filepath"
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
		initCommandHandler()
	case "config":
		fmt.Println("gate says Hello")
	case "run":
		fmt.Println("gate says Hello")
	default:
		return fmt.Errorf("unknown or invalid gate command: %s\nSee `gate help` to see commands", subcommand)
	}

	return nil

}

func initCommandHandler() error {

	osConfigPath, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config directory: %w", err)
	}

	gateConfigDirPath := filepath.Join(osConfigPath, "gate-cli")

	_, err = os.Stat(gateConfigDirPath)
	if err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check gate config directory: %w", err)
	}

	err = os.MkdirAll(gateConfigDirPath, 0700)
	if err != nil {
		return fmt.Errorf("failed to create gate config directory: %w", err)
	}

	sampleConfigFilePath := filepath.Join(gateConfigDirPath, "sampleConfigFile.json")
	sampleConfigFile, err := os.OpenFile(sampleConfigFilePath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)

	if err != nil {
		if os.IsExist(err) {
			return nil
		}

		return fmt.Errorf("failed to create config file: %w", err)
	}

	defer sampleConfigFile.Close()

	defaultConfig := []byte(`{
		"steps":[]
	}`)

	_, err = sampleConfigFile.Write(defaultConfig)
	if err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
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
