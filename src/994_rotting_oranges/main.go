package main

type Cell struct{ r, c int }

func orangesRotting(grid [][]int) int {
	queue := []Cell{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, Cell{i, j})
			}
		}
	}

	minute := -1
	for len(queue) > 0 {
		size := len(queue)
		dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				nr, nc := current.r+dir[0], current.c+dir[1]
				if nc == -1 || nr == -1 || nr == len(grid) || nc == len(grid[0]) || grid[nr][nc] == 2 || grid[nr][nc] == 0 {
					continue
				}
				grid[nr][nc] = 2
				queue = append(queue, Cell{nr, nc})
			}
		}
		minute++
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return max(0, minute)

}
