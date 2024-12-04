package day01

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	input := strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3`)

	score := part2(input)

	assert.Equal(t, 31, score)
}
