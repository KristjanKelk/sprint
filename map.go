package sprint

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result
}

func IsNegative(n int) bool {
	return n < 0
}

func IsPrime(n int) bool {
	// Check if n is less than or equal to 1
	if n <= 1 {
		return false
	}

	// Check if n is 2 or 3
	if n == 2 || n == 3 {
		return true
	}

	// Check if the sum of digits is divisible by 3
	sumOfDigits := 0
	for temp := n; temp > 0; temp /= 10 {
		sumOfDigits += temp % 10
	}
	if sumOfDigits%3 == 0 {
		return false
	}

	// Check if n is of the form 6k ± 1
	if n%6 != 1 && n%6 != 5 {
		return false
	}

	// Check for divisors up to the square root of n
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
