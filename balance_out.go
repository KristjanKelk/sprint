package sprint

func BalanceOut(arr []bool) []bool {
	count := 0
	for _, value := range arr {
		if value == true {
			count++
		}
	}

	// Calculate the number of excess true or false values
	excess := count - len(arr)/2

	// Create a new slice with balanced values
	balancedSlice := []bool{}
	for _, value := range arr {
		if value == true && excess > 0 {
			excess--
		} else {
			balancedSlice = append(balancedSlice, value)
		}
	}

	return balancedSlice
}
