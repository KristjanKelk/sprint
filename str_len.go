package sprint

import "unicode/utf8"

func StrLen(s string) []int {
	result := []int{utf8.RuneCountInString(s), len(s)}
	return result
}
