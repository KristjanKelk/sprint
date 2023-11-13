package sprint

func IterativeFactorial(n int) int {
	factorial := 1

	if n == 0 {
		return 0
	}

	if n < 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial
}
