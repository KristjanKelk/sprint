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

/*chatgpt pakutud compact versioon minu koodist
package sprint

func BalancedParentheses(input string) bool {
	var leftParent, rightParent int
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case ']', '}', ')':
			rightParent++
		case '[', '{', '(':
			leftParent++
		}
	}
	return leftParent == rightParent
}
*/
