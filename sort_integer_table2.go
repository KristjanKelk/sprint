package sprint

func SortIntegerTable2(table []int) []int {
	sortedSlice := make([]int, 0, len(table)) // Initialize with capacity
	startingPoint := 0

	for i := 0; i < len(table); i++ {
		if table[i] < startingPoint {
			startingPoint = table[i]
		}
		sortedSlice = append(sortedSlice, table[i]) // Append the current element to sortedSlice
	}

	for i := 0; i < len(sortedSlice)-1; i++ { // Use len(sortedSlice)-1 to avoid index out of range
		for j := i + 1; j < len(sortedSlice); j++ {
			if sortedSlice[i] > sortedSlice[j] {
				sortedSlice[i], sortedSlice[j] = sortedSlice[j], sortedSlice[i]
			}
		}
	}

	return sortedSlice
}
