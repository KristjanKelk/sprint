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

	if countTrue > countFalse {
		for i := 0; i < countTrue-countFalse; i++ {
			arr = append(arr, false)
		}
	} else if countTrue < countFalse {
		for i := 0; i < countFalse-countTrue; i++ {
			arr = append(arr, true)
		}
	}

	return arr
}
