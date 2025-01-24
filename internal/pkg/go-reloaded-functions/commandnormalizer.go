package goreloaded

import (
	"regexp"
	"strings"
)

// CommandNormalize scans the input string and trims unnecessary spaces inside commands
func CommandNormalize(input string) string {
	// Regular expression to normalize commands of the form (word , number)
	commandRe := regexp.MustCompile(`\(\s*([^,]+?)\s*,\s*([^,]+?)\s*\)`)
	normalized := commandRe.ReplaceAllStringFunc(input, func(match string) string {
		// Normalize strings starting and ending with parentheses
		inner := strings.Trim(match, "()")
		parts := strings.SplitN(inner, ",", 2)
		if len(parts) == 2 {
			cmd := strings.TrimSpace(parts[0])
			count := strings.TrimSpace(parts[1])
			return "(" + cmd + ", " + count + ")"
		}
		return match
	})

	// Separate punctuation that follows normalized parentheses blocks
	spacePunctuationRe := regexp.MustCompile(`\)([.,!?])`)
	normalized = spacePunctuationRe.ReplaceAllString(normalized, ") $1")

	// General regex to handle any non-space character followed by parentheses
	wordParenthesesRe := regexp.MustCompile(`(\S)\s*\(([^,]+?)\)`)
	normalized = wordParenthesesRe.ReplaceAllString(normalized, "$1 ($2)")

	return normalized
}
