package main

func minCostClimbingStairs(cost []int) int {
	// [1,100,1,1,1,100,1,1,100,1]
	// db[i] is min cost to get to i
	// db[0] = 0
	// db[1] = 0
	// db[2] = min(cost[0], cost[1]) => 1
	// db[3] = min(cost[1], cost[2]) => 1
	// db[i] = min(cost[i-1], cost[i-2]) => min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
	dp := []int{0, 0}

	for i := 2; i <= len(cost); i++ {
		dp = append(dp, min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]))
	}

	return dp[len(cost)]
}
