package sprint

func GCD(a, b int) int {
	result := 1
	min := a
	if min > b {
		min = b
	} else if a == b {
		result = a
		return result
	} else if a == 0 || b == 0 {
		return 0
	}

	for i := min / 2; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			result = i
			return result
		}
	}
	return result
}
