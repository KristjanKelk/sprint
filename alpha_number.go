package sprint

func AlphaNumber(n int) string {
	//if n equals 0, write 'a'
	if n == 0 {
		return "a"
	}
	//declaring variables
	result := ""
	negative := false

	//if input is negative, set variable negative true
	if n < 0 {
		negative = true
		n = -n
	}

	//using modulus to get the digits
	for n > 0 {
		digit := n % 10
		letter := rune('a' + digit)
		result = string(letter) + result
		n /= 10
	}

	//if negative is true add - in front of result

	if negative {
		result = "-" + result
	}

	return result
}
