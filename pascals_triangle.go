package sprint

func PascalsTriangle(n int) [][]int {
	pascTri := [][]int{}
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = pascTri[i-1][j-1] + pascTri[i-1][j]
		}
		pascTri = append(pascTri, row)
	}

	return pascTri
}
