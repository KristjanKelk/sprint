package sprint

func Payout(amount int, denominations []int) (payout []int) {
	if amount == 0 {
		return payout
	}

	bubbleSort(denominations)

	// Store the original amount for comparison later
	originalAmount := amount

	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}

	// Use the stored amount for the comparison
	if originalAmount == summ(payout) {
		return payout
	}

	return []int{}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func summ(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
