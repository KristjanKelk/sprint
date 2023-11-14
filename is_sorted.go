package sprint

func IsSorted(arr []string, descending bool) bool {
	n := len(arr)

	// Define the comparison function based on the desired order
	compare := func(a, b string) int {
		if descending {
			return len(b) - len(a)
		}
		return len(a) - len(b)
	}

	// Iterate over the elements of the slice and compare adjacent pairs
	for i := 0; i < n-1; i++ {
		result := compare(arr[i], arr[i+1])

		// Check if the order is correct based on the comparison result
		if result > 0 {
			return false
		} else if result < 0 {
			return false
		}
	}

	return true
}
