package cmd

import (
	"github.com/arichardsmith/aoc-2024-go/challenge/day01"
	"github.com/spf13/cobra"
)

func NewApp() *cobra.Command {
	app := &cobra.Command{
		Use:     "aoc2024",
		Short:   "Advent of Code 2024",
		Example: "aoc2024 day01 [args]",
	}

	day01.RegisterCommand(app)

	return app
}
