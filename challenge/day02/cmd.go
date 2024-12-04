package day02

import (
	"fmt"

	"github.com/arichardsmith/aoc-2024-go/internal"
)

func RunPart(part int) error {
	input := internal.MustHaveInputFile()

	switch part {
	case 1:
		_, err := fmt.Println(part1(input))
		return err
	case 2:
		_, err := fmt.Println(part2(input))
		return err
	}

	return fmt.Errorf("no part %d", part)
}
