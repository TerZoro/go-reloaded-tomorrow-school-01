package goreloaded

import (
	"strconv"
	"strings"
	"unicode"
)

func ApplyCommands(input string) string {
	words := strings.Fields(input)
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]
		if isCommand(word) {
			cmd, count := parseCommand(&word)
			if i+1 < len(words) && (strings.HasSuffix(words[i+1], ")")) {
				count = parseNumber(&words[i+1])
			}
			applyTransformation(&result, cmd, count)
			continue
		}
		if word != "" {
			result = append(result, word)
		}

	}

	// Rebuild the final string from the transformed words
	return strings.Join(result, " ")
}

func isCommand(word string) bool {
	return strings.Contains(word, "(")
}

func parseCommand(command *string) (string, int) {
	inner := strings.TrimSpace(strings.Trim(*command, "()"))
	parts := strings.Split(inner, ",")

	cmd := strings.TrimSpace(strings.Trim(parts[0], ","))
	count := 1
	if len(parts) == 2 {
		num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err == nil {
			count = num
		}
	}
	*command = cmd

	return cmd, count
}

func parseNumber(word *string) int {
	inner := strings.TrimRight(*word, "(),.!?;:")
	num, err := strconv.Atoi(inner)
	if err != nil {
		return 1
	}
	*word = ""
	// Remove the processed number
	return num
}

func applyTransformation(result *[]string, cmd string, count int) {
	r := *result
	if len(r) == 0 {
		return
	}

	if count > len(r) {
		count = len(r)
	}

	start := len(r) - 1
	for i := start; i >= 0; i-- {
		if count == 0 {
			break
		}
		if !isNotWord(r[i]) {
			switch cmd {
			case "low":
				r[i] = Low(r[i])
			case "up":
				r[i] = Up(r[i])
			case "cap":
				r[i] = Cap(r[i])
			case "hex":
				r[i] = Hex(r[i])
			case "bin":
				r[i] = Bin(r[i])
			}
			count--
		}
	}
	*result = r
}
func isNotWord(word string) bool {
	for _, r := range word {
		if unicode.IsPunct(r) {
			return true
		}
	}
	return false
}
