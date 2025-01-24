// cmd/app/main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	helper "goreloaded/internal/pkg/helper"
)

func main() {
	// Validate command-line arguments
	inputFile, outputFile, err := helper.ValidateArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Define the fixed directory
	samplesDir := "test/testdata/samples"
	resultsDir := "test/testdata/results"

	// Ensure the directory exists
	if err := os.MkdirAll(samplesDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating samples directory %s: %v\n", samplesDir, err)
		os.Exit(1)
	}
	if err := os.MkdirAll(resultsDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating results directory %s: %v\n", resultsDir, err)
		os.Exit(1)
	}

	// Construct the full path
	inputPath := filepath.Join(samplesDir, inputFile)
	outputPath := filepath.Join(resultsDir, outputFile)

	// Read the input file
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	text := string(data)
	lines := strings.Split(text, "\n")
	results := []string{}

	// Process each line
	for _, line := range lines {
		processedLine, err := helper.ErrorProcessing(line)
		if err != nil {
			fmt.Printf("Error processing line: %v\n", err)
			continue
		}

		results = append(results, processedLine)
	}

	// Join the results
	outputText := JoinLines(results)

	// Write to the output file
	if err := os.WriteFile(outputPath, []byte(outputText), 0644); err != nil {
		fmt.Printf("Error writing to file %s: %v\n", outputPath, err)
		os.Exit(1)
	}

	fmt.Println("Processed text saved to:", outputPath)
}

func JoinLines(lines []string) string {
	return strings.Join(lines, "\n")
}
