package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	excludeSlice := []float64{}

	if from < 0 {
		from = 0
	}
	if from >= len(arr) {
		from = len(arr)
	}
	if to < 0 {
		to = 0
	}
	if to >= len(arr) {
		to = len(arr)
	}

	if from > to {
		from, to = to, from
	}

	for i := 0; i < len(arr); i++ {
		if i < from || i >= to {
			excludeSlice = append(excludeSlice, arr[i])
		}
	}

	return excludeSlice
}
