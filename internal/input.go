package internal

import (
	"bufio"
	"io"
	"iter"
	"os"
	"path/filepath"
	"runtime"
)

// InputFile returns an io.Reader for the file pointed at by the --input flag. If the flag is not specified, it looks
// for a file named "input.txt" in the same package as the caller.
// Borrowed from https://github.com/nlowe/aoc2023/blob/master/challenge/input.go
func DefaultInputFile() io.Reader {
	_, f, _, ok := runtime.Caller(1)
	if !ok {
		panic("failed to determine input path, provide it with -i instead")
	}

	path := filepath.Join(filepath.Dir(f), "input.txt")

	r, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return r
}

// Returns an iterator over the lines in the given reader.
func Lines(r io.Reader) iter.Seq[string] {
	scanner := bufio.NewScanner(r)

	return func(yield func(string) bool) {
		for scanner.Scan() {
			if err := scanner.Err(); err != nil && err != io.EOF {
				panic(err)
			}

			if !yield(scanner.Text()) {
				return
			}
		}
	}
}
