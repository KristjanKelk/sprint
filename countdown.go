package sprint

import "fmt"

func Countdown(n int) string {
	countdown := ""
	for i := n; i >= 1; i -= 2 {
		countdown += fmt.Sprintf("%d, ", i)
	}
	return countdown + "0!"
}
