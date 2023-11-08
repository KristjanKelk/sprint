package sprint

import (
	"strings"
)

func AtoiBase(s string, base string) int {
	// Step 1: Validate the base
	charMap := make(map[rune]bool)
	for _, char := range base {
		if charMap[char] == true || char == '+' || char == '-' {
			return 0
		}
		charMap[char] = true
	}
	if len(charMap) < 2 {
		return 0
	}

	result := 0
	baseLength := len(base)

	// Process each character in the input string in reverse order
	for _, char := range s {
		// Find the index of the character in the base string
		index := strings.IndexRune(base, char)
		result = result*baseLength + index
	}

	return result
}
