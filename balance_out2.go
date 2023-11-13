package sprint

func BalanceOut2(arr []bool) []bool {
	falses := 0
	result := arr
	truth := 0

	for _, value := range arr {
		if value == true {
			truth++
		} else {
			falses++
		}
	}

	if falses == truth {
		return result
	}

	if truth > falses {
		for i := 0; i < truth-falses; i++ {
			arr = append(arr, false)
		}
	} else if truth < falses {
		for i := 0; i < falses-truth; i++ {
			arr = append(arr, true)
		}
	}
	return arr
}
