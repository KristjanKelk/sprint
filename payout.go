package sprint

func Payout(amount int, denominations []int) (payout []int) {
	bubbleSort(denominations)

	if amount == 0 {
		return []int{}
	}

	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}
	if payout == nil {
		return []int{}
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
