package sprint

func Payout(amount int, denominations []int) (payout []int) {
	bubbleSort(denominations)

	for _, denom := range denominations {
		for amount >= denom {
			payout = append(payout, denom)
			amount -= denom
		}
	}

	return payout
}

func Sum(p []int) int {
	sum := 0
	for i := 0; i < len(p); i++ {
		sum += p[i]
	}
	return sum
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
