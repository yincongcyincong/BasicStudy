package util

import (
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	fd, err := os.Open(path)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(fd)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}