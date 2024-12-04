package day02

import (
	"fmt"
	"io"
)

func part1(input io.Reader) int {
	safe := 0

	for report := range parseFile(input) {
		if err := validateReport(report); err == nil {
			safe++
		}
	}

	return safe
}

const (
	INCREASE = "increase"
	DECREASE = "decrease"
)

func validateReport(report Report) error {
	var v Validator

	n := len(report)

	for i := 1; i < n; i++ {
		if err := v.Validate(report[i-1], report[i]); err != nil {
			return err
		}
	}

	return nil
}

type Validator struct {
	dir string
}

func (v *Validator) Validate(a, b int) error {
	var p_dir string

	if a > b {
		p_dir = DECREASE
	} else {
		p_dir = INCREASE
	}

	if v.dir == "" {
		v.dir = p_dir
	} else if v.dir != p_dir {
		return fmt.Errorf("direction is %s, was %s", p_dir, v.dir)
	}

	delta := abs(b - a)

	if delta > 3 || delta < 1 {
		return fmt.Errorf("delta %d (%d - %d) out of range", delta, b, a)
	}

	return nil
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}
