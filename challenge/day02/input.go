package day02

import (
	"io"
	"iter"
	"strconv"
	"strings"

	"github.com/arichardsmith/aoc-2024-go/internal"
)

type Report = []int

func parseFile(r io.Reader) iter.Seq[Report] {
	return func(yield func(Report) bool) {

		for line := range internal.Lines(r) {
			res := make([]int, 0)
			for _, v := range strings.Split(line, " ") {
				p, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}

				res = append(res, p)
			}

			yield(res)
		}

	}
}
