package main

func pacificAtlantic(heights [][]int) [][]int {
	pacific, atlantis := make([][]bool, len(heights)), make([][]bool, len(heights))
	for i := range pacific {
		pacific[i] = make([]bool, len(heights[i]))
	}
	for i := range atlantis {
		atlantis[i] = make([]bool, len(heights[i]))
	}

	var dfs func(i, j, previous int, visited [][]bool)

	dfs = func(i, j, previous int, visited [][]bool) {
		if i == -1 || j == -1 || i == len(heights) || j == len(heights[0]) ||
			visited[i][j] || previous > heights[i][j] {
			return
		}

		visited[i][j] = true

		dfs(i-1, j, heights[i][j], visited)
		dfs(i+1, j, heights[i][j], visited)
		dfs(i, j-1, heights[i][j], visited)
		dfs(i, j+1, heights[i][j], visited)

	}

	for i := range len(heights) {
		dfs(i, 0, 0, pacific)
		dfs(i, len(heights[0])-1, 0, atlantis)
	}

	for j := range len(heights[0]) {
		dfs(0, j, 0, pacific)
		dfs(len(heights)-1, j, 0, atlantis)
	}

	result := [][]int{}
	for i := range heights {
		for j := range heights[i] {
			if pacific[i][j] && atlantis[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
