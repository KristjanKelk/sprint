package sprint

func Accumulate(n int) int {
	if n < 0 {
		return 0
	} else {
		sum := 0
		for i := 0; i <= n; i++ {
			sum += i
		}
		return sum
	}
}

/*

func Accumulate(n int) int {
	if n < 0 {
		return 0
	}
	return Accumulate(n-1) + n
}




*/
