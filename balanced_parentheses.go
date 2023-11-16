package sprint

func BalancedParentheses(input string) bool {
	var leftParent int
	var rightParent int
	for i := 0; i < len(input); i++ {
		if input[i] == ']' || input[i] == '}' || input[i] == ')' {
			rightParent += 1
		} else if input[i] == '[' || input[i] == '{' || input[i] == '(' {
			leftParent += 1
		}
	}
	if leftParent == rightParent {
		return true
	}
	return false
}
