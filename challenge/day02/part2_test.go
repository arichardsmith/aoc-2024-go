package day02

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	// Test against the example input
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	expected := 4

	res := part2(strings.NewReader(input))

	assert.Equal(t, expected, res)
}

func TestDampDelta(t *testing.T) {
	assert.True(t, validateNext([]int{1, 2, 10, 4, 5}))
	assert.False(t, validateNext([]int{1, 2, 10, 9, 10, 11}))
}

func TestDampDirection(t *testing.T) {
	assert.True(t, validateNext([]int{1, 2, 3, 4, 5}))
	assert.False(t, validateNext([]int{1, 2, 4, 2, 3, 4}))
}

func TestRemovingLast(t *testing.T) {
	assert.True(t, validateNext([]int{1, 2, 3, 4, 5, 10}))
	assert.True(t, validateNext([]int{1, 2, 3, 4, 5, 4}))
}

func TestRemovingFirst(t *testing.T) {
	assert.True(t, validateNext([]int{1, 5, 6, 7, 8, 9}))
	assert.False(t, validateNext([]int{5, 2, 3, 4, 5, 2}))
	assert.True(t, validateNext([]int{6, 10, 9, 8, 7, 6}))
}
