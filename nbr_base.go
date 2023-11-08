package main

func NbrBase(n int, base string) string {
	// Check if base contains at least 2 unique characters and does not contain '+' or '-'
	charMap := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return "NV"
		}
		charMap[char] = true
	}
	if len(charMap) < 2 {
		return "NV"
	}

	// Convert the integer to the specified base
	var result string
	if n < 0 {
		result = "-"
		n = -n
	}
	for n > 0 {
		remainder := n % len(base)
		result = string(base[remainder]) + result
		n = n / len(base)
	}

	if result == "" {
		result = "0"
	}
	return result
}
