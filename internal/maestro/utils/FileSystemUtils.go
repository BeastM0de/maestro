package utils

import (
	"fmt"
	"os"
)

func EnsurePathExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to check if directory exists: %w", err)
	}
	return nil
}
