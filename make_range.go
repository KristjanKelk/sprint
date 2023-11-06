package sprint

func MakeRange(min, max int) []int {

	if min >= max {
		return nil
	} else {
		rangeSlice := []int{}
		for i := min; i < max; i++ {
			rangeSlice = append(rangeSlice, i)
		}
		return rangeSlice

	}
}
