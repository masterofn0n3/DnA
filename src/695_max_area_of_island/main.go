package main

func maxAreaOfIsland(grid [][]int) int {

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == -1 || j == -1 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0

		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)

	}
	maxA := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			maxA = max(maxA, dfs(i, j))
		}
	}
	return maxA
}
