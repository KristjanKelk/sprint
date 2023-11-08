package sprint

import (
	"strings"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {

	// Step 1: Validate the base
	for _, char := range nbr {
		if strings.IndexRune(baseFrom, char) == -1 {
			return "NV"
		}
		// Rest of the conversion logic
	}
	number := 0
	baseLength := len(baseFrom)

	// Process each character in the input string in reverse order
	for _, char := range nbr {
		// Find the index of the character in the base string
		index := strings.IndexRune(baseFrom, char)
		number = number*baseLength + index
	}

	// Convert the integer to the specified base
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
