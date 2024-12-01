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

	res := calculateDistance(strings.NewReader(input))

	assert.Equal(t, 11, res)
}
