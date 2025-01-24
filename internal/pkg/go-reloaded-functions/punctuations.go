package goreloaded

import (
	"strings"
	"unicode"
)

func Punctuation(input string) string {
	var result strings.Builder
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if isPunctuationChar(char) {
			// Remove space before the punctuation
			if result.Len() > 0 && unicode.IsSpace(rune(result.String()[result.Len()-1])) {
				trimLastSpace(&result)
			}
			result.WriteRune(char)

			// Add a space after punctuation if the next character is not punctuation or end of string
			if i+1 < len(runes) && !isPunctuationChar(runes[i+1]) && !unicode.IsSpace(runes[i+1]) {
				result.WriteRune(' ')
			}
			continue
		}
		result.WriteRune(char)
	}

	return result.String()
}

func isPunctuationChar(char rune) bool {
	return strings.ContainsRune(".,!?;:", char)
}

func trimLastSpace(sb *strings.Builder) {
	str := sb.String()
	str = strings.TrimRight(str, " ")
	sb.Reset()
	sb.WriteString(str)
}
