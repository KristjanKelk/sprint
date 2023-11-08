package sprint

import "strings"

func NbrBase(n int, base string) string {
	if !isValidBase(base) {
		return "NV"
	}

	if n == 0 {
		return "0"
	}

	isNegative := n < 0
	if isNegative {
		n = -n
	}

	result := ""
	baseLength := len(base)

	for n > 0 {
		digit := n % baseLength
		result = string(base[digit]) + result
		n /= baseLength
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i, ch := range base {
		if ch == '+' || ch == '-' || strings.Count(base, string(ch)) > 1 {
			return false
		}

		for _, otherCh := range base[i+1:] {
			if ch == otherCh {
				return false
			}
		}
	}

	return true
}
