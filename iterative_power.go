package sprint

func IterativePower(n int, power int) int {
	if power < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= power; i++ {
		result *= n
	}
	return result
}
