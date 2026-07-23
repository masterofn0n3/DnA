package main

import "math"

func maxProfit(prices []int) int {
	//hold[i] represent max profit if we buy on day i
	hold := []int{-prices[0]}
	sold := []int{math.MinInt}
	rest := []int{0}

	for i := 1; i < len(prices); i++ {
		hold = append(hold, max(hold[i-1], rest[i-1]-prices[i]))
		sold = append(sold, hold[i-1]+prices[i])
		rest = append(rest, max(rest[i-1], sold[i-1]))
	}

	return max(sold[len(sold)-1], rest[len(rest)-1])

}
