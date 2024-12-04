package cmd

import (
	"fmt"

	"github.com/arichardsmith/aoc-2024-go/challenge/day01"
	"github.com/arichardsmith/aoc-2024-go/challenge/day02"
	"github.com/spf13/cobra"
)

func NewApp() *cobra.Command {
	app := &cobra.Command{
		Use:     "aoc2024",
		Short:   "Advent of Code 2024",
		Example: "aoc2024 day01 [args]",
		Run:     run,
	}

	app.PersistentFlags().IntP("day", "d", 0, "Day of advent (1-25)")
	app.PersistentFlags().IntP("part", "p", 1, "Part of the challenge (1-2)")
	app.MarkPersistentFlagRequired("day")

	registerCreateCommand(app)

	return app
}

func run(cmd *cobra.Command, args []string) {
	day, _ := cmd.Flags().GetInt("day")
	part, _ := cmd.Flags().GetInt("part")

	if part <= 0 {
		fmt.Println("Part must be >= 1")

		cmd.Help()
		return
	}

	var err error

	switch day {
	case 1:
		err = day01.RunPart(part)
	case 2:
		err = day02.RunPart(part)
	default:
		fmt.Printf("No day %d\n\n", day)
		fmt.Printf("Run `aoc2024 create --day %d` to create a new day\n", day)
	}

	if err != nil {
		fmt.Printf("Error running part %d: %s\n", part, err)
	}
}
