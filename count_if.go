package sprint

func CountIf(f func(string) bool, tab []string) int {
	result := []bool{}
	for _, v := range tab {
		if f(v) == true {
			result = append(result, f(v))
		} else {
			continue
		}
	}
	return len(result)
}

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, char := range s {
		if char < 'A' || char > 'Z' {
			return false
		}

	}
	return true
}

func IsAlphanumeric(s string) bool {
	for _, char := range s {
		switch {
		case '0' <= char && char <= '9':
		case 'A' <= char && char <= 'Z':
		case 'a' <= char && char <= 'z':
			continue
		default:
			return false
		}
	}
	return true
}
