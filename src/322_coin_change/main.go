package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	// dp[i] represent amount of count to make up to i
	for i := range dp {
		for _, c := range coins {
			if c > i {
				continue
			}
			dp[i] = min(dp[i-c]+1, dp[i])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
