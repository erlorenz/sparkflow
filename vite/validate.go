package vite

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/erlorenz/sparkflow"
)

func parseExt(logicalPath string) (string, error) {
	ext := filepath.Ext(logicalPath)

	if !slices.Contains([]string{".js", ".ts", ".css"}, ext) {
		return "", fmt.Errorf("%w: %s", sparkflow.ErrInvalidExt, ext)
	}
	return ext, nil
}

func parseEnvironment(environment string) string {
	if environment == "development" {
		return environment
	}
	return "production"
}
