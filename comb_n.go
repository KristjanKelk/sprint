package sprint

import "fmt"

func CombN(n int) []string {
	if n < 1 || n > 9 {
		return nil
	}

	var combinations []string
	digits := make([]int, n)

	generateCombinations(0, n, digits, &combinations)
	return combinations
}

func generateCombinations(index, n int, digits []int, combinations *[]string) {
	if index == n {
		combination := ""
		for i := 0; i < n; i++ {
			combination += fmt.Sprintf("%d", digits[i])
		}
		*combinations = append(*combinations, combination)
		return
	}

	start := 0
	if index > 0 {
		start = digits[index-1] + 1
	}

	for i := start; i <= 9; i++ {
		digits[index] = i
		generateCombinations(index+1, n, digits, combinations)
	}
}
