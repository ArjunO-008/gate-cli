package cmd

import "fmt"

const GATE_VERSION string = "v1.0.0"

func ShowGateVersion() {
	fmt.Printf("gate version %s\n", GATE_VERSION)
}
