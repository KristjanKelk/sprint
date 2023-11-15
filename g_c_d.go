package sprint

func GCD(a, b int) int {
	min := a
	if min < b {
	} else {
		min = b
	}
	result := 1
	for i := min / 2; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			result = i
			return result
		}
	}
	return result
}
