package day01

import (
	"github.com/spf13/cobra"
)

func RegisterCommand(app *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "day01",
		Short: "Day 1 Solutions",
	}

	p1 := &cobra.Command{
		Use:   "part01",
		Short: "Part 1 Solution",
		Run:   part1,
	}

	p2 := &cobra.Command{
		Use:   "part02",
		Short: "Part 2 Solution",
		Run:   part2,
	}

	cmd.AddCommand(p1, p2)

	app.AddCommand(cmd)
}
