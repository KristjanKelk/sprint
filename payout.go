package sprint

func Payout(amount int, denominations []int) (payout []int) {
	// Sort denominations in ascending order
	bubbleSort(denominations)

	for i := len(denominations) - 1; i >= 0; i-- {
		if denominations[i] <= amount {
			payout = append(payout, denominations[i])
			amount -= denominations[i]
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
			if arr[j] > arr[j+1] {
				// Swap elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
