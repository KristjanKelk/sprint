package sprint

func StrLen(s string) []int {
	result := []int{characterCount(s), len(s)}
	return result
}
func characterCount(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
