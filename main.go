package main

import (
	"os"

	"github.com/arichardsmith/aoc-2024-go/cmd"
)

func main() {
	if err := cmd.NewApp().Execute(); err != nil {
		os.Exit(1)
	}
}
