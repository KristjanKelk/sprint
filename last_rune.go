package sprint

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	lastRune := rune(s[len(s)-1])
	return lastRune
}
