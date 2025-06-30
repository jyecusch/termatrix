package main

import (
	"os"

	"github.com/jyecusch/termatrix/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
