package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		inDegree[v[0]]++
	}

	queue := []int{}

	for i, v := range inDegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	result := []int{}

	for len(queue) > 0 {
		current := queue[0]
		result = append(result, current)
		queue = queue[1:]
		for _, v := range graph[current] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}

	}

	if len(result) == numCourses {
		return result
	}

	return []int{}
}
