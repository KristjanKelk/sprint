package sprint

func BasicAtoi(s string) int {

	newInteger := 0
	//controlling if the digit is between 0 -9
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		newInteger = newInteger*10 + digit
	}

	return newInteger
}
