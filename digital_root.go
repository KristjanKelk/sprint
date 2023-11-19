package sprint

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	sum := 0

	for n != 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}

	return DigitalRoot(sum)
}
