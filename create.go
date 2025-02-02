package gofiles

import "os"

func CreateWriteFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
}

func CreateDirectoryIfNotExist(path string, perm os.FileMode) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, perm); err != nil {
			return err
		}
	}

	return nil
}
