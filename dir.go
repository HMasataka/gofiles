package gofiles

import (
	"os"
	"path/filepath"
)

func Pwd() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
