package sprint

func BalanceOut(arr []bool) []bool {
	countTrue := 0
	countFalse := 0

	// Count the number of true and false values in the input slice
	for _, value := range arr {
		if value == true {
			countTrue++
		} else {
			countFalse++
		}
	}

	// Calculate the number of true and false values needed to balance the slice
	balanceCount := min(countTrue, countFalse)

	// Create a new slice with balanced true and false values
	balancedSlice := make([]bool, 2*balanceCount)

	for i := 0; i < balanceCount; i++ {
		balancedSlice[i] = true
		balancedSlice[i+balanceCount] = false
	}

	return balancedSlice
}
