package sprint

func StrRev(s string) string {
	revStr := []rune(s) // Convert the input string to a slice of runes
	length := len(revStr)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		revStr[i], revStr[j] = revStr[j], revStr[i]
	}

	return string(revStr)
}
