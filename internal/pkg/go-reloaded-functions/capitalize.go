package goreloaded

import (
	"strings"
	"unicode"
)

func Cap(str string) string {
	if str == "" {
		return ""
	}
	runes := []rune(strings.ToLower(str))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
