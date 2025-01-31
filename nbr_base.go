package sprint

// NbrBase(-125, "choumi")
func NbrBase(n int, base string) string {
	// Step 1: Validate the base
	charMap := make(map[rune]bool)
	for _, char := range base {
		if charMap[char] == true || char == '+' || char == '-' {
			return "NV"
		}
		charMap[char] = true
	}
	if len(charMap) < 2 {
		return "NV"
	}

	// Step 2: Convert the integer to the specified base
	var result string
	if n < 0 {
		n = -n
		for n > 0 {
			remainder := n % len(base)
			result = string(base[remainder]) + result
			n = n / len(base)
		}
		result = "-" + result
	} else {
		for n > 0 {
			remainder := n % len(base)
			result = string(base[remainder]) + result
			n = n / len(base)
		}
	}

	// Step 3: Handle negative numbers and return the result
	if result == "" {
		result = "0"
	}
	return result
}
