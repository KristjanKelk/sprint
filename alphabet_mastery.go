package sprint

func AlphabetMastery(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		char := 'a' + i
		result += string(char)
	}
	return result
}
