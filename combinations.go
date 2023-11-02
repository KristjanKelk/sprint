package sprint

import "fmt"

func Combinations() string {
	combinations := ""
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {
				combinations += fmt.Sprintf("%d%d%d", i, j, k)

			}
		}
	}
	if len(combinations) > 0 {
		combinations = combinations[:len(combinations)-2]
	}

	return combinations

}
