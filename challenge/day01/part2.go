package day01

import (
	"io"
)

func part2(input io.Reader) int {
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
