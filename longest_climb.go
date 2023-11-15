package sprint

func LongestClimb(arr []int) []int {
	result := []int{}
	tempArr := []int{arr[0]}

	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			tempArr = append(tempArr, arr[i])
		} else {
			if len(result) < len(tempArr) {
				result = tempArr
			}
			tempArr = []int{arr[i]}
		}
	}

	// Check if the last climb is the longest
	if len(result) < len(tempArr) {
		result = tempArr
	}

	return result
}
