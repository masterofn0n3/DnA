package main

func longestIncreasingPath(matrix [][]int) int {

	memo := make(map[[2]int]int)

	var dfs func(i, j, previous int) int

	dfs = func(i, j, previous int) int {
		if i == -1 || i == len(matrix) || j == -1 || j == len(matrix[0]) || matrix[i][j] <= previous {
			return 0
		}

		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		current := matrix[i][j]
		count := max(dfs(i+1, j, current), dfs(i-1, j, current), dfs(i, j+1, current), dfs(i, j-1, current))

		memo[[2]int{i, j}] = count + 1

		return count + 1
	}

	maxCount := 0
	for i := range matrix {
		for j := range matrix[i] {
			maxCount = max(maxCount, dfs(i, j, -1))
		}
	}

	return maxCount
}
