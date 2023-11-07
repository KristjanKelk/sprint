package sprint

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) < 'a' || rune(s[i]) > 'z' {
			return false
		}
	}
	return true
}
