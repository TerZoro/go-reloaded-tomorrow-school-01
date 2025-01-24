package helper

import (
	"errors"
	goreloaded "goreloaded/internal/pkg/go-reloaded-functions"
	"path/filepath"
	"strings"
)

func ValidateArgs(args []string) (string, string, error) {
	if len(args) != 3 {
		return "", "", errors.New("usage: go run ./cmd/app/ <sample*.txt> <result*.txt>")
	}

	inputFile := args[1]
	outputFile := args[2]

	// Ensure input file starts with "sample" and ends with ".txt"
	if !strings.HasPrefix(inputFile, "sample") || !strings.HasSuffix(inputFile, ".txt") {
		return "", "", errors.New("input file must start with 'sample' and have a .txt extension")
	}

	// Ensure output file starts with "result" and ends with ".txt"
	if !strings.HasPrefix(outputFile, "result") || !strings.HasSuffix(outputFile, ".txt") {
		return "", "", errors.New("output file must start with 'result' and have a .txt extension")
	}

	// Ensure outputFile does not contain any directory paths
	if filepath.Base(outputFile) != outputFile {
		return "", "", errors.New("output file must not contain directory paths")
	}

	return inputFile, outputFile, nil
}

func ErrorProcessing(line string) (string, error) {
	trimmedLine := strings.TrimSpace(line)
	if trimmedLine == "" {
		return "", nil
	}

	processed := goreloaded.CommandNormalize(trimmedLine)
	processed = goreloaded.ApplyCommands(processed)
	processed = goreloaded.Punctuation(processed)
	processed = goreloaded.Apostrophe(processed)
	processed = goreloaded.Articles(processed)

	// If any processing step can fail, handle errors. For now, we assume they are infallible
	return processed, nil
}
