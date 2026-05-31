package main

func exist(board [][]byte, word string) bool {

	var dfs func(index, row, col int) bool

	dfs = func(index, row, col int) bool {
		if index == len(word) {
			return true
		}
		if row < 0 || col < 0 || row > len(board)-1 || col > len(board[0])-1 {
			return false
		}
		char := board[row][col]
		if char != word[index] || char == '#' {
			return false
		}
		board[row][col] = '#'
		found := dfs(index+1, row+1, col) ||
			dfs(index+1, row-1, col) ||
			dfs(index+1, row, col+1) ||
			dfs(index+1, row, col-1)
		board[row][col] = char

		return found
	}

	for row := range board {
		for col := range board[row] {
			if dfs(0, row, col) {
				return true
			}
		}
	}

	return false
}
