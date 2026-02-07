package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var gateVersion string = "v0.0.1"

type Config struct {
	Version  string   `json:"version"`
	Name     string   `json:"name"`
	Settings Settings `json:"settings"`
	Steps    []Step   `json:"steps"`
	Extras   Extras   `json:"extras"`
}

type Settings struct {
	WorkingDirectory string `json:"workingDirectory"`
}

type Step struct {
	Type    string `json:"type"`
	Command string `json:"command"`
	Args    string `json:"args,omitempty"`
	Dir     string `json:"dir,omitempty"`
}

type Extras struct {
	Git GitExtras `json:"git"`
}

type GitExtras struct {
	Enabled bool   `json:"enabled"`
	Config  string `json:"config"`
}

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
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to check gate config directory: %w", err)
		}

		err = os.MkdirAll(gateConfigDirPath, 0700)
		if err != nil {
			return fmt.Errorf("failed to create gate config directory: %w", err)
		}

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

	config := Config{
		Version: "0.1.0",
		Name:    "sampleconfig",
		Settings: Settings{
			WorkingDirectory: ".",
		},
		Steps: []Step{
			{
				Type:    "executable",
				Command: "npm",
				Args:    "install",
				Dir:     "project",
			},
			{
				Type:    "shellexecutable",
				Command: "echo Build completed successfully",
			},
		},
		Extras: Extras{
			Git: GitExtras{
				Enabled: false,
				Config:  "gitConfig",
			},
		},
	}

	encoder := json.NewEncoder(sampleConfigFile)
	encoder.SetIndent("", " ")
	err = encoder.Encode(config)

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
