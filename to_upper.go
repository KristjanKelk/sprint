package sprint

func ToUpper(s string) string {
	result := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		if char >= 'a' && char <= 'z' {
			char -= 'a' - 'A'
		}
		result[i] = char
	}
	return string(result)
}
