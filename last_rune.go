package sprint

func LastRune(s string) rune {
	lastRune := rune(s[0])
	for i := 0; i <= len(s); i++ {
		lastRune = rune(s[i])
	}
	return lastRune
}
