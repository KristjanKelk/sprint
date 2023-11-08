package sprint

import (
	"strings"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {

	// Step 1: Validate the base
	charMap := make(map[rune]bool)
	for _, char := range baseFrom {
		if charMap[char] == true || char == '+' || char == '-' {
			return "NV"
		}
		charMap[char] = true
	}
	if len(charMap) < 2 {
		return "NV"
	}

	number := 0
	baseLength := len(baseFrom)

	// Process each character in the input string in reverse order
	for _, char := range nbr {
		// Find the index of the character in the base string
		index := strings.IndexRune(baseFrom, char)
		number = number*baseLength + index
	}

	// Step 1: Validate the base

	// Step 2: Convert the integer to the specified base
	var result string
	if number < 0 {
		number = -number
		for number > 0 {
			remainder := number % len(baseTo)
			result = string(baseTo[remainder]) + result
			number = number / len(baseTo)
		}
		result = "-" + result
	} else {
		for number > 0 {
			remainder := number % len(baseTo)
			result = string(baseTo[remainder]) + result
			number = number / len(baseTo)
		}
	}

	// Step 3: Handle negative numbers and return the result
	if result == "" {
		result = "0"
	}
	return result

}
