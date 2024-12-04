package day02

import (
	"io"
)

func part2(input io.Reader) int {
	safe := 0

	for report := range parseFile(input) {
		if validateNext(report) {
			safe++
		}
	}

	return safe
}

func validateNext(report Report) bool {
	if err := validateReport(report); err == nil {
		return true
	}

	for i := 0; i < len(report); i++ {
		t := make(Report, len(report))
		copy(t, report)

		if err := validateReport(append(t[:i], t[i+1:]...)); err == nil {
			return true
		}
	}

	return false
}
