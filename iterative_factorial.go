package sprint

func IterativeFactorial(n int) int {
	isNegative := false
	factorial := 1

	if n == 0 {
		return 0
	}

	if n < 0 {
		isNegative = true
		n = -n
	}

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	if isNegative {
		factorial *= -1
	}

	return factorial
}
