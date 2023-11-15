package sprint

func RemoveDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	var uniqueArr []int

	for _, num := range arr {
		// If the element is not in the map, it's unique
		if !seen[num] {
			seen[num] = true
			uniqueArr = append(uniqueArr, num)
		}
	}

	return uniqueArr
}
