package gofiles

import (
	"bufio"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadFileYield(path string) (func(func(string) bool), error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return func(yield func(string) bool) {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			yield(scanner.Text())
		}
	}, nil
}
