package sprint

func Split(s, sep string) []string {
	words := []string{}
	word := ""
	separatorPos := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep[separatorPos] {
			separatorPos++
			if separatorPos == len(sep) {
				// Found the full separator sequence
				if word != "" {
					words = append(words, word)
					word = ""
				}
				separatorPos = 0
			}
		} else {
			if separatorPos > 0 {
				// Append any partially matched separator runes to the word
				word += sep[:separatorPos]
				separatorPos = 0
			}
			word += string(s[i])
		}
	}
	if word != "" {
		words = append(words, word)
	}

	return words
}
