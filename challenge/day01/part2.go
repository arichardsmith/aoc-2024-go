package day01

import (
	"fmt"
	"io"

	"github.com/arichardsmith/aoc-2024-go/internal"
	"github.com/spf13/cobra"
)

func part2(cmd *cobra.Command, args []string) {
	fmt.Println(calculateSimilarityScore(internal.DefaultInputFile()))
}

func calculateSimilarityScore(input io.Reader) int {
	m := make(map[int]int)

	var left []int

	// Initialize the map with the left list values.
	for pair := range parseFile(input) {
		m[pair.Right]++
		left = append(left, pair.Left)
	}

	total := 0

	for _, v := range left {
		total += v * m[v]
	}

	return total
}
