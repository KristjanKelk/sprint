package sprint

func NRune(s string, i int) rune {
	if len(s) == 0 || i >= len(s) {
		return 0
	}

	return rune(s[i])
}
