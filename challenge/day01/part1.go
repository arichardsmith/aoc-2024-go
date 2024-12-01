package day01

import (
	"io"
	"math"
	"slices"
)

func part1(input io.Reader) int {
	n := 0

	var left, right sortedSlice
	for pair := range parseFile(input) {
		left.Insert(pair.Left)
		right.Insert(pair.Right)
		n++
	}

	total := 0

	for i := 0; i < n; i++ {
		total += int(math.Abs(float64(left.values[i] - right.values[i])))
	}

	return total
}

type sortedSlice struct {
	values []int
}

func (s *sortedSlice) Insert(val int) {
	if len(s.values) == 0 {
		s.values = append(s.values, val)

		return
	}

	for i, v := range s.values {
		if v >= val {
			if i == 0 {
				s.values = slices.Insert(s.values, 0, val)
			} else {
				s.values = slices.Insert(s.values, i, val)
			}

			return
		}
	}

	s.values = append(s.values, val)
}
