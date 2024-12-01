package internal

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
	"path/filepath"
	"runtime"
)

// InputFile returns an io.Reader for the file pointed at by the --input flag. If the flag is not specified, it looks
// for a file named "input.txt" in the same package as the caller.
// Borrowed from https://github.com/nlowe/aoc2023/blob/master/challenge/input.go
func DefaultInputFile() (io.Reader, error) {
	path, err := HereN(2, "input.txt")
	if err != nil {
		return nil, err
	}

	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func MustHaveInputFile() io.Reader {
	r, err := DefaultInputFile()
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

// Get a path relative to the caller's file.
func Here(elem ...string) (string, error) {
	return HereN(2, elem...)
}

func HereN(n int, elem ...string) (string, error) {
	_, f, _, ok := runtime.Caller(n)
	if !ok {
		return "", fmt.Errorf("failed to determine call file")
	}

	return filepath.Join(filepath.Dir(f), filepath.Join(elem...)), nil
}
