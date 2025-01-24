package goreloaded

import "strings"

func IsTxt(input string) bool {
	s := strings.ToLower(input)
	return strings.HasSuffix(s, ".txt")
}
