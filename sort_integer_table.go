package sprint

func SortIntegerTable(table []int) []int {
	max := table[0]
	result := []int{}
	for i := 0; i < len(table); i++ {
		if max < table[i] {
			max = table[i]
		}
	}

	for len(table) > 0 {
		min := 0
		for i := 0; i < len(table); i++ {
			if table[i] < table[min] {
				min = i
			}

		}
		result = append(result, table[min])
		table = append(table[:min], table[min+1:]...)

	}

	return result
}
