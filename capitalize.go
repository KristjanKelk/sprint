package sprint

func Capitalize(s string) string {
	capitalize := true
	result := make([]rune, len(s))

	for i, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if capitalize {
				if char >= 'a' && char <= 'z' {
					char -= 'a' - 'A'
				}
				capitalize = false
			}
			result[i] = char
		} else {
			capitalize = true
			result[i] = char
		}
	}

	return string(result)
}
