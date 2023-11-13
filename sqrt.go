package sprint

func Sqrt(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return n
	}

	squareRoot := n / 2 //first try

	for squareRoot*squareRoot > n {
		squareRoot = (squareRoot + n/squareRoot) / 2
	}

	if squareRoot*squareRoot == n {
		return squareRoot
	}

	return 0
}
