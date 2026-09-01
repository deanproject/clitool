package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	slog.Info("Hello")
	readDir()
}

func readDir() {
	dirPath := "./logfinder"

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		slog.Error("failed to read directory", "path", dirPath, "error", err)
		return
	}

	for _, fname := range entries {
		if fname.IsDir() {
			slog.Info("found a directory", "name", fname.Name())
		} else {
			slog.Info("found a file", "name", fname.Name())
		}
	}
}
