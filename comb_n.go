package sprint

import "fmt"

// CombN generates combinations of n digits from 1 to 9 and returns them as a slice of strings.
func CombN(n int) []string {
	// Check if the input 'n' is valid (between 1 and 9).
	if n < 1 || n > 9 {
		return nil // Return nil for invalid input values.
	}

	// Initialize an empty slice to store the generated combinations.
	var combinations []string

	// Initialize a slice to hold the digits for each combination.
	digits := make([]int, n)

	// Generate combinations recursively and update the 'combinations' slice.
	generateCombinations(0, n, digits, &combinations)

	// Return the generated combinations.
	return combinations
}

// generateCombinations is a recursive function to generate combinations of digits.
func generateCombinations(index, n int, digits []int, combinations *[]string) {
	// If 'index' reaches 'n', a complete combination is generated.
	if index == n {
		// Construct a combination string by joining the digits.
		combination := ""
		for i := 0; i < n; i++ {
			combination += fmt.Sprintf("%d", digits[i])
		}

		// Append the combination to the 'combinations' slice.
		*combinations = append(*combinations, combination)
		return
	}

	// Determine the starting digit for the current position.
	start := 0
	if index > 0 {
		start = digits[index-1] + 1
	}

	// Iterate through possible digits for the current position.
	for i := start; i <= 9; i++ {
		// Set the current position in 'digits' to the chosen digit.
		digits[index] = i

		// Recursively generate combinations for the remaining positions.
		generateCombinations(index+1, n, digits, combinations)
	}
}
