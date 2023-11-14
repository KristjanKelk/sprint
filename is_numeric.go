package sprint

func isNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) < '0' || rune(s[i]) > '9' {
			return false
		}
	}
	return true
}
