package sprint

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	// Ensure from and to are within valid bounds
	if from < 0 {
		from = 0
	}
	if from > len(arr) {
		from = len(arr)
	}
	if to < 0 {
		to = 0
	}
	if to > len(arr) {
		to = len(arr)
	}

	// Ensure from is less than or equal to to
	if from > to {
		from, to = to, from
	}

	// Create a new slice with elements outside the specified range
	modifiedArray := append(arr[:from], arr[to:]...)

	return modifiedArray
}
