package sprint

func BulkAtoi(arr []string) []int {
	result := []int{}
	for _, val := range arr {
		result = append(result, atoi(val))
	}
	return result
}

// ATOI
func atoi(s string) int {
	newInteger := 0
	sign := 1
	// Does same thing but with a one if function

	if len(s) > 0 {
		if s[0] == '-' {
			sign = -1
			s = s[1:] // Remove the leading '-'
		} else if s[0] == '+' {
			s = s[1:] // Remove the leading '+'
		}
	}

	// Controlling if the digit is between 0-9
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		newInteger = newInteger*10 + digit
	}

	return newInteger * sign
}
