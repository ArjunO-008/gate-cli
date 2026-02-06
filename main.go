package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Println("gate is Active")
		os.Exit(0)
	} else {
		fmt.Fprintln(os.Stderr, "Use: gate <command>")
		os.Exit(2)
	}
}

func listCommands() {
	fmt.Printf("")
}
