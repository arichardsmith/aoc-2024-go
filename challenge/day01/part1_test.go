package day01

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	i := 1
	input := `3   10
4   20
5   30
`
	for pair := range parseFile(strings.NewReader(input)) {
		assert.Equal(t, i+2, pair.Left)
		assert.Equal(t, i*10, pair.Right)

		i++
	}
}

func TestCalculateDistance(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3
`

	res := part1(strings.NewReader(input))

	assert.Equal(t, 11, res)
}

func TestSortedSlice(t *testing.T) {
	s := sortedSlice{}

	s.Insert(1)
	assert.Equal(t, []int{1}, s.values)

	s.Insert(3)
	s.Insert(2)
	s.Insert(4)
	s.Insert(2)
	assert.Equal(t, []int{1, 2, 2, 3, 4}, s.values)
}
