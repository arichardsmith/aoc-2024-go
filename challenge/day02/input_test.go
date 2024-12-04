package day02

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	// Test against the example input
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	var res []Report

	for report := range parseFile(strings.NewReader(input)) {
		res = append(res, report)
	}

	assert.Equal(t,
		[]Report{
			[]int{7, 6, 4, 2, 1},
			[]int{1, 2, 7, 8, 9},
			[]int{9, 7, 6, 2, 1},
			[]int{1, 3, 2, 4, 5},
			[]int{8, 6, 4, 4, 1},
			[]int{1, 3, 6, 7, 9},
		},
		res)
}
