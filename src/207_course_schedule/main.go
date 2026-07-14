package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseGraph := make([][]int, numCourses)

	states := make([]int, numCourses)

	for i := range prerequisites {
		courseGraph[prerequisites[i][0]] = append(courseGraph[prerequisites[i][0]], prerequisites[i][1])
	}

	var dfs func(course int) bool

	dfs = func(course int) bool {
		if states[course] == 1 {
			return false
		}
		if states[course] == 2 {
			return true
		}
		states[course] = 1

		for _, v := range courseGraph[course] {
			if !dfs(v) {
				return false
			}
		}

		states[course] = 2

		return true
	}

	for i := range courseGraph {
		if !dfs(i) {
			return false
		}
	}

	return true

}
