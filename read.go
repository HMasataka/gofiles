package gofiles

import "os"

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
