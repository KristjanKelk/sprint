package sprint

func SplitWhitespaces(s string) []string {
	words := []string{}
	word := ""

	for _, char := range s {
		if isSeparator(char) {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}

func isSeparator(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}
