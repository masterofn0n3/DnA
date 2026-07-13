package main

func solve(board [][]byte) {

	var dfs func(i, j int)

	dfs = func(i, j int) {
		if i == -1 || j == -1 || i == len(board) || j == len(board[0]) || board[i][j] == 'X' || board[i][j] == '#' {
			return
		}

		board[i][j] = '#'

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := range len(board) {
		dfs(i, 0)
		dfs(i, len(board[0])-1)
	}

	for j := range len(board[0]) {
		dfs(0, j)
		dfs(len(board)-1, j)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

}
