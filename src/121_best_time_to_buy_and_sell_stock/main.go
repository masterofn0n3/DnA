package main

func maxProfit(prices []int) int {
	hold, sold := make([]int, len(prices)), make([]int, len(prices))
	hold[0] = -prices[0]
	sold[0] = 0

	// 2 4 1
	// hold sold
	// -2 0
	// -2 2
	// -1
	for i := 1; i < len(prices); i++ {
		hold[i] = max(hold[i-1], -prices[i])
		sold[i] = max(sold[i-1], hold[i-1]+prices[i])
	}

	return max(sold[len(sold)-1], 0)

}
