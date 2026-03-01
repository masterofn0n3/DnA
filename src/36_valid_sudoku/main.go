package validsudoku

//I need to check 3 things
//Horizontal rows, which is the inner array
//Vertical rows, which is the outter one
// which can be done with 1 loop through every array
// And the grid

/*
1 2 3 * * * * * 9
* * 4 * * * * * 2
6 * * * * * * * *

	i = 0 =>
	0-0 0-1 0-2 0-3 0-4 0-5
	1-0 1-1 1-2 1-3 1-4 1-5
	2-0 2-1 2-2 2-3 2-4 2-5
*/
func isValidSudoku(board [][]byte) bool {
	var rowOcc [9][9]bool
	var colOcc [9][9]bool
	var boxOcc [9][9]bool

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			d := int(board[i][j] - '1')
			boxIdx := (i/3)*3 + j/3
			if rowOcc[i][d] || colOcc[j][d] || boxOcc[boxIdx][d] {
				return false
			}
			rowOcc[i][d] = true
			colOcc[j][d] = true
			boxOcc[boxIdx][d] = true

		}
	}
	return true
}
