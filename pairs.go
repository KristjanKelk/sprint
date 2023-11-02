package sprint

import "fmt"

func Pairs() string {
	combinations := ""
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			combinations += fmt.Sprintf("%02d %02d, ", i, j)
		}
	}

	if len(combinations) > 0 {
		combinations = combinations[:len(combinations)-2]
	}
	return combinations
}
