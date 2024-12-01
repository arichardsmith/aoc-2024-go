package day01

import (
	"github.com/arichardsmith/aoc-2024-go/internal"
	"github.com/spf13/cobra"
)

func RegisterCommand(app *cobra.Command) {
	cmd := internal.NewCommand("day01", "Day 1 Solution", run)
	app.AddCommand(cmd)
}

func run(cmd *cobra.Command, args []string) {

}
