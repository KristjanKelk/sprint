package sprint

import (
	"strconv"
	"strings"
)

func EightQueens() string {
	var result string
	queenColumns := make([]int, 8)
	placeQueens(queenColumns, 0, &result)
	return strings.TrimSpace(result)
}

func placeQueens(queenColumns []int, row int, result *string) {
	if row == 8 {
		*result = (*result + "\n" + queensToString(queenColumns))
		return
	}

	for col := 0; col < 8; col++ {
		if isSafe(queenColumns, row, col) {
			queenColumns[row] = col
			placeQueens(queenColumns, row+1, result)
		}
	}
}

func isSafe(queenColumns []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if queenColumns[i] == col || // Check column
			queenColumns[i] == col-(row-i) || // Check left diagonal
			queenColumns[i] == col+(row-i) { // Check right diagonal
			return false
		}
	}
	return true
}

func queensToString(queenColumns []int) string {
	var result strings.Builder
	for _, col := range queenColumns {
		result.WriteString(strconv.Itoa(col + 1)) // Convert index to string
	}
	return result.String()
}
