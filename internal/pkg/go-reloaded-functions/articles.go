package goreloaded

import (
	"strings"
	"unicode"
)

func Articles(input string) string {
	var result strings.Builder
	words := strings.Fields(input) // Split input into words (handles spacing naturally)

	for i, word := range words {
		if (strings.EqualFold(word, "a") || strings.EqualFold(word, "A")) && i+1 < len(words) {
			nextWord := words[i+1]
			if StartsWithVowelOrH(nextWord) {
				if strings.EqualFold(word, "a") {
					result.WriteString("an ")
					continue
				}
				result.WriteString("An ")
			} else {
				if strings.EqualFold(word, "a") {
					result.WriteString("a ")
					continue
				}
				result.WriteString("A ")

			}
			continue
		}
		result.WriteString(word)

		// Add space after the word, except for the last one
		if i < len(words)-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

func StartsWithVowelOrH(s string) bool {
	if s == "" {
		return false
	}
	r := unicode.ToLower(rune(s[0]))
	return strings.ContainsRune("aeiouh", r)
}
