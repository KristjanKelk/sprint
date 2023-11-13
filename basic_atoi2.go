package sprint

func BasicAtoi2(s string) int {

	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			result = result*10 + int(rune(s[i]-'0'))
		} else {
			return 0
		}

	}
	return result
}
