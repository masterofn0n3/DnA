package main

func findTargetSumWays(nums []int, target int) int {

	memo := make(map[int]map[int]int)
	var dfs func(sum, i int) int
	dfs = func(sum, i int) int {
		if i >= len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		if v, ok := memo[i][sum]; ok {
			return v
		}

		count := dfs(sum-nums[i], i+1) + dfs(sum+nums[i], i+1)
		if memo[i] == nil {
			memo[i] = map[int]int{}
		}
		memo[i][sum] = count

		return count
	}

	return dfs(0, 0)

}
