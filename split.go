package sprint

func Split(s, sep string) []string {
	words := []string{}
	word := ""
	separatorMatched := 0
	if sep == "" {
		return nil
	}

	for i := 0; i < len(s); i++ {
		if s[i] == sep[separatorMatched] {
			separatorMatched++
			if separatorMatched == len(sep) {
				// Found the full separator sequence
				separatorMatched = 0
				words = append(words, word)
				word = ""
			}
		} else {
			// Append any partially matched separator runes to the word
			word += sep[:separatorMatched]
			separatorMatched = 0
			word += string(s[i])
		}
	}

	// Append any remaining characters to the last word
	word += sep[:separatorMatched]
	words = append(words, word)

	return words
}
