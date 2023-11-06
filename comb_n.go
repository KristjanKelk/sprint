package sprint

import (
	"fmt"
	"math"
)

func CombN(n int) []string {
	if n < 1 || n > 9 {
		return nil
	}

	var combinations []string
	maxDigit := 9
	digits := make([]int, n)
	base := int(math.Pow10(n - 1))

	generateCombinations(0, 0, n, maxDigit, base, digits, &combinations)
	return combinations
}

func generateCombinations(index, start, n, maxDigit, base int, digits []int, combinations *[]string) {
	if index == n {
		combination := ""
		for i := 0; i < n; i++ {
			combination += fmt.Sprintf("%d", digits[i])
		}
		*combinations = append(*combinations, combination)
		return
	}

	for i := start; i <= maxDigit; i++ {
		digits[index] = i
		generateCombinations(index+1, i+1, n, maxDigit, base, digits, combinations)
	}
}
