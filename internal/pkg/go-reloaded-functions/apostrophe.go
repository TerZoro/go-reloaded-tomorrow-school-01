package goreloaded

import (
	"strings"
)

func Apostrophe(input string) string {
	var result strings.Builder
	inQuotes := false
	start := 0

	for i, char := range input {
		if char == '\'' {
			if inQuotes {
				result.WriteString(strings.TrimSpace(input[start:i]))
				result.WriteString("'")
				inQuotes = false
			} else {
				result.WriteString("'")
				start = i + 1
				for start < len(input) && input[start] == ' ' {
					start++
				}
				inQuotes = true
			}
		} else if !inQuotes {
			result.WriteRune(char)
		}
	}

	return result.String()
}
