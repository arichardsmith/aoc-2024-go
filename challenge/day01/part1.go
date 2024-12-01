package day01

import (
	"io"
	"math"
	"sort"
)

func part1(input io.Reader) int {
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
