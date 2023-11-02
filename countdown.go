package sprint

func Countdown(n int) string {
	countdown := ""
	for i := n; i > 0; i -= 2 {
		char := '0' + rune(i)
		countdown += string(char)
		countdown += ", "
	}
	return countdown + "0!"
}
