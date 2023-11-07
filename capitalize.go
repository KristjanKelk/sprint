package sprint

func Capitalize(s string) string {
	capitalize := true
	result := make([]rune, len(s))

	for i, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if capitalize {
				if char >= 'a' && char <= 'z' {
					char -= 'a' - 'A'
				} else if char >= '0' && char <= '9' {
					char = char
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
