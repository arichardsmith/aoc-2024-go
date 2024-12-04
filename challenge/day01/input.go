package day01

import (
	"io"
	"iter"
	"strconv"
	"strings"

	"github.com/arichardsmith/aoc-2024-go/internal"
)

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
