package main

// nums = [1,2,3,1]
func rob(nums []int) int {
	// dp[i] = max ammount can rob
	// i = 0 dp = 1
	// i = 1 dp = 2, can not robb adjacent
	// i = 3 dp = 4, max(dp[i-2]+cost[i], dp[i-1])
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := []int{nums[0], max(nums[1], nums[0])}

	for i := 2; i <= len(nums); i++ {
		dp = append(dp, max(dp[i-2]+nums[i], dp[i-1]))
	}

	return dp[len(nums)-1]
}
