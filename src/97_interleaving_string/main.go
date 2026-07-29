package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	cache := make(map[int]map[int]bool)

	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		if i+j == len(s3) {
			return true
		}

		if v, ok := cache[i][j]; ok {
			return v
		}

		result := false
		if i < len(s1) && s1[i] == s3[i+j] {
			result = result || dfs(i+1, j)
		}
		if j < len(s2) && s2[j] == s3[i+j] {
			result = result || dfs(i, j+1)
		}

		if _, ok := cache[i]; !ok {
			cache[i] = make(map[int]bool)
		}
		cache[i][j] = result

		return result

	}

	return dfs(0, 0)

}
