package main

import "math"

func solveNQueens(n int) [][]string {
	result := [][]string{}

	queens := make([]int, n)

	var place func(row int)

	place = func(row int) {
		if row == n {
			result = append(result, convert(queens, n))
			return
		}

		for col := range n {
			if isValid(row, col, queens) {
				queens[row] = col
				place(row + 1)
			}
		}
	}

	place(0)

	return result
}

func convert(queens []int, n int) []string {
	result := []string{}
	for _, v := range queens {
		row := make([]byte, n)
		for i := range row {
			row[i] = '.'
		}
		row[v] = 'Q'
		result = append(result, string(row))
	}
	return result

}

func isValid(row, col int, queens []int) bool {
	for r := range row {
		qCol := queens[r]

		if qCol == col {
			return false
		}

		if math.Abs(float64(row-r)) == math.Abs(float64(col-qCol)) {
			return false
		}
	}
	return true
}
