package sprint

func RecursiveFactorial(n int) int {
	// Base case: factorial of 0 is 1
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	return n * RecursiveFactorial(n-1)
}
