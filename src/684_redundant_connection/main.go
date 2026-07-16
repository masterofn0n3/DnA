package main

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)

	for i := range parent {
		parent[i] = i
	}

	var find func(i int) int

	find = func(i int) int {
		current := i
		for current != parent[current] {
			current = parent[current]
		}
		return current
	}

	for _, v := range edges {
		if find(v[0]) == find(v[1]) {
			return v
		}
		parent[find(v[0])] = find(v[1])
	}

	return nil
}
