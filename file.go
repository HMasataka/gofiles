package gofiles

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateWriteFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
}

func CopyFile(src string, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("can't stat %s: %w", src, err)
	}

	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("can't copy non-regular source file %s (%s)", src, srcInfo.Mode().String())
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("can't open source file %s: %w", src, err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return fmt.Errorf("can't make destination directory %s: %w", filepath.Dir(dst), err)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("can't create destination file %s: %w", dst, err)
	}
	defer dstFile.Close()

	size, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("can't copy data: %w", err)
	}
	if size != srcInfo.Size() {
		return fmt.Errorf("incomplete copy, %d of %d", size, srcInfo.Size())
	}

	return dstFile.Sync()
}

func WriteLines(path string, lines []string) error {
	file, err := CreateWriteFile(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}

func ReadFileAll(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadLinesYield(path string) (func(func(string) bool), error) {
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
