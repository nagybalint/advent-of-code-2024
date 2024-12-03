package utils

import (
	"log"
	"os"
	"path/filepath"
)

// Reads a file from a path relative to the project root
func ReadFileFromRelative(relativePath string) (string, error) {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Println("Cannot determine working directory")
		return "", nil
	}
	buf, err := os.ReadFile(filepath.Join(projectRoot, relativePath))
	if err != nil {
		log.Println("Cannot read contents of file", err)
		return "", err
	}

	return string(buf), nil
}
