package sprint

func Payout(amount int, denominations []int) (payout []int) {

	if amount == 0 {
		return payout
	}

	// If the requested amount is not equal to the sum of denominations, return an empty slice.
	if amount != summ(denominations) {
		return []int{}
	}

	bubbleSort(denominations)
	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}
	return payout
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
