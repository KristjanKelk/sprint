package sprint

func SortWordArr(a []string) []string {
	max := a[0]
	result := []string{}
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	for len(a) > 0 {
		min := 0
		for i := 0; i < len(a); i++ {
			if a[i] < a[min] {
				min = i
			}

		}
		result = append(result, a[min])
		a = append(a[:min], a[min+1:]...)

	}

	return result
}
