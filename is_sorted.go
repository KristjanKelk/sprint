package sprint

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	b := a
	result := make([]string, len(a))
	copy(result, a) // Copy the original slice to avoid modifying it

	// Sorting the slice using the provided comparison function
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if f(result[i], result[j]) > 0 {
				// Swap elements if they are in the wrong order
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	if b == result {
		return true
	} else {
		return false
	}
}

func Compare(a, b string) int {
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	for i := 0; i < minLength; i++ {
		diff := int(a[i]) - int(b[i])
		if diff != 0 {
			return diff
		}
	}

	return len(a) - len(b)
}
