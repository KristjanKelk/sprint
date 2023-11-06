package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	excludeSlice := []float64{}

	if from > to {
		for i := 0; i < len(arr); i++ {
			if i < to || i >= from {
				excludeSlice = append(excludeSlice, arr[i])
			}
		}
	} else {
		for i := 0; i < len(arr); i++ {
			if i < from || i >= to {
				excludeSlice = append(excludeSlice, arr[i])
			}
		}
	}

	return excludeSlice
}
