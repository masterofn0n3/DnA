package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1
	for _, coin := range coins {
		for i := 1; i < len(dp); i++ {
			if coin > i {
				continue
			}
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
