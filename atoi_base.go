package sprint

import (
	"strings"
)

func AtoiBase(s string, base string) int {
	// Step 1: Validate the base
	charMap := make(map[rune]bool) // Create a map to keep track of characters in the base
	for _, char := range base {    // Iterate through each character in the base
		if charMap[char] == true || char == '+' || char == '-' { // Check if the character is already seen or is '+' or '-'
			return 0
		}
		charMap[char] = true // Mark the character as seen
	}
	if len(charMap) < 2 { // Check if there are fewer than 2 unique characters in the base
		return 0 // Return 0 to indicate an error
	}

	result := 0             // Initialize a variable to store the result
	baseLength := len(base) // Calculate the length of the base

	// Process each character in the input string in reverse order
	for _, char := range s { // Iterate through each character in the input string

		index := strings.IndexRune(base, char) // Find the index of the character in the base string
		result = result*baseLength + index     // Update the result based on the character index
	}

	return result
}
