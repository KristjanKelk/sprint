package sprint

import "fmt"

func StrCompress(input string) string {
	result := ""
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			result += fmt.Sprintf("%d%c", count, input[i-1])
			count = 1
		}
	}

	result += fmt.Sprintf("%d%c", count, input[len(input)-1])

	return result
}
