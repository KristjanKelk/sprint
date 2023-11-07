package sprint

// Capitalize function capitalizes the first letter of each word while preserving spaces and non-alphanumeric characters.
func Capitalize(s string) string {
	result := make([]rune, len(s))
	capitalizeNext := true // A flag to indicate whether the next character should be capitalized
	//if no input
	if len(s) < 1 {
		return s
	}

	for i, char := range s {
		if isAlphaNumeric(char) { // Check if the character is alphanumeric
			if capitalizeNext {
				result[i] = toUpper(char) // Capitalize the first letter of the word
				capitalizeNext = false
			} else {
				result[i] = toLower(char) // Convert the rest of the word to lowercase
			}
		} else {
			// Non-alphanumeric character, set flag to capitalize the next letter
			capitalizeNext = true
			result[i] = char
		}

	}
	return string(result)
}

// isAlphaNumeric checks if a character is alphanumeric (a letter or a digit).
func isAlphaNumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

// toUpper converts a lowercase character to uppercase.
func toUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}

// toLower converts an uppercase character to lowercase.
func toLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char - 'A' + 'a'
	}
	return char
}
