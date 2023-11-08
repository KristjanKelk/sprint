package sprint

func Split(s, sep string) []string {
	words := []string{}
	word := ""
	separatorSet := make(map[rune]bool)

	for _, char := range sep {
		separatorSet[char] = true
	}

	for _, char := range s {
		if separatorSet[char] {
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
