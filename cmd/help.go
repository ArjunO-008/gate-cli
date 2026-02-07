package cmd

import "fmt"

func DefineHelpCommand() {
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
