package sprint

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {

		match := true

		for j := 0; j < len(toFind); j++ {

			if toFind[j] != s[i+j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}
	return -1
}
