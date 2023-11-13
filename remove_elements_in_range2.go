package sprint

func RemoveElementsInRange2(arr []float64, from, to int) []float64 {
	if from > to {
		from, to = to, from
	}
	if from < 0 {
		from = 0
	} else if to > len(arr) {
		to = len(arr)
	}

	modiArr := append(arr[:from], arr[to:]...)
	return modiArr
}
