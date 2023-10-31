package sprint

func ReverseAlphabetValue(ch rune) rune {
	reversing := rune('z' - (ch - 'a'))
	return reversing
}
