package cmd

import "fmt"

const GATE_VERSION string = "v0.1.0"

func ShowGateVersion() {
	fmt.Printf("gate version %s\n", GATE_VERSION)
}
