package main

func ladderLength(beginWord string, endWord string, wordList []string) int {

	var find func(s string) []string

	find = func(s string) []string {
		result := []string{}
		for _, v := range wordList {
			diffCount := 0
			for i := range v {
				if v[i] != s[i] {
					diffCount++
				}
				if diffCount > 1 {
					break
				}
			}
			if diffCount == 1 {
				result = append(result, v)
			}

		}
		return result
	}

	queue := []string{beginWord}
	visited := map[string]struct{}{}
	count := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == endWord {
				return count + 1
			}
			if _, ok := visited[current]; ok {
				continue
			}
			visited[current] = struct{}{}
			relates := find(current)
			for _, v := range relates {
				if _, ok := visited[v]; !ok {
					queue = append(queue, v)
				}
			}
		}
		count++
	}

	return 0

}
