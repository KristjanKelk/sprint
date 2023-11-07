package sprint

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]

	for i := 1; i < len(strs); i++ {
		result += sep + strs[i] // Concatenate the separator and the next string
	}
	return result

}
