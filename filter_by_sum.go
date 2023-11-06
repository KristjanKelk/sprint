func FilterBySum(arr [][]int, limit int) [][]int {
	result := [][]int{}
	sum := 0
	for _, subArr := range arr {
		for _, val := range subArr {
			sum += val
		}
		if sum >= limit {
			result = append(result, subArr)
		}
		sum = 0

	}
	return result
}