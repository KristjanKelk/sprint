package sprint

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(n int) int {
	if n == 0 {
		return 5
	}
	if n%2 == 0 {
		n++
	}

	for !isPrime(n) {
		n += 2
	}

	return n
}
