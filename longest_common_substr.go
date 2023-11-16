package sprint

func LongestCommonSubstr(str1, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}

	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	var result string
	var tempResult string
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if i < len(str1)-1 && j < len(str2)-1 && str1[i] == str2[j] && str1[i+1] == str2[j+1] {
				tempResult = tempResult + string(str2[j])
				i++
			} else if str1[i] == str2[j] && (i == len(str1)-1 || j == len(str2)-1 || str1[i+1] != str2[j+1]) {
				tempResult = tempResult + string(str2[j])
				if len(tempResult) > len(result) {
					result = tempResult
				}
				tempResult = ""
			}
		}
	}

	if len(tempResult) > len(result) {
		result = tempResult
	}

	return result
}
