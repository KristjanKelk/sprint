package sprint

func MakeRange(min, max int) []int {

	var nilSlice []int
	if min >= max {
		return nilSlice
	} else {
		rangeSlice := []int{}
		for i := min; min < max; i++ {
			rangeSlice = append(rangeSlice, i)
		}
		return rangeSlice

	}
}
