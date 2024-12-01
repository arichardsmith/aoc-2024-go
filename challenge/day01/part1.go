package day01

import (
	"fmt"
	"io"
	"iter"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/arichardsmith/aoc-2024-go/internal"
	"github.com/spf13/cobra"
)

func part1(cmd *cobra.Command, args []string) {
	fmt.Println(calculateDistance(internal.DefaultInputFile()))
}

func calculateDistance(input io.Reader) int {
	n := 0

	var left, right []int
	for pair := range parseFile(input) {
		left = append(left, pair.Left)
		right = append(right, pair.Right)
		n++
	}

	// TODO: Insert the numbers in the correct order when reading.
	sort.Ints(left)
	sort.Ints(right)

	total := 0

	for i := 0; i < n; i++ {
		total += int(math.Abs(float64(left[i] - right[i])))
	}

	return total
}

type ListItem struct {
	Left  int
	Right int
}

func parseFile(file io.Reader) iter.Seq[ListItem] {
	return func(yield func(ListItem) bool) {
		for line := range internal.Lines(file) {
			parts := strings.Fields(line)

			if len(parts) != 2 {
				panic("invalid line format")
			}

			left, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}

			right, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			yield(ListItem{
				Left:  left,
				Right: right,
			})
		}
	}
}
