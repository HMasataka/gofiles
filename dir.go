package gofiles

import (
	"os"
	"path/filepath"
)

func Pwd() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func ListFiles(path string) ([]string, error) {
	var filePaths []string

	if err := filepath.Walk(path, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.IsDir() {
			return nil
		}

		filePaths = append(filePaths, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return filePaths, nil
}

func ListFilesYield(path string) (func(func(string) bool), error) {
	filePaths, err := ListFiles(path)
	if err != nil {
		return nil, err
	}

	return func(yield func(string) bool) {
		for _, filePath := range filePaths {
			yield(filePath)
		}
	}, nil
}
