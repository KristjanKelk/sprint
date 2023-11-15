package sprint

func GCD(a, b int) int {
	result := 1
	min := a
	if min > b && b != 0 {
		min = b
	} else if a == b && a != 0 {
		result = a
		return result
	} else if a == 0 || b == 0 {
		return a + b
	}

	for i := min / 2; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			result = i
			return result
		}
	}
	return -1
}
