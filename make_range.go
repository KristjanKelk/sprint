package sprint

func MakeRange(min, max int) []int {

	var nilSlice []int
	if min >= max {
		return nilSlice
	} else {
		for i := min; min < max; i++ {
			rangeSlice := []int{}
			rangeSlice = append(rangeSlice, i)
		}
	}
}
