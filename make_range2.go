package sprint

func MakeRange2(min, max int) []int {
	if min > max {
		min, max = max, min
	}
	result := []int{}
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
