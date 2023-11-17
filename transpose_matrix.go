package sprint

func TransposeMatrix(matrix [][]int) [][]int {
	// Get the number of rows and columns in the original matrix
	rows := len(matrix)
	cols := len(matrix[0])

	// Create a new matrix with swapped rows and columns
	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	// Iterate through the original matrixand fill the transposed matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
