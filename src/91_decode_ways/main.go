package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	dp := []int{1, 1}
	for i := 2; i <= len(s); i++ {
		dp = append(dp, 0)
		if isInRange(string(s[i-1 : i])) {
			dp[i] += dp[i-1]
		}
		if isInRange(s[i-2 : i]) {
			dp[i] += dp[i-2]
		}

	}

	return dp[len(s)]
}

func isInRange(s string) bool {
	if s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	return num > 0 && num <= 26
}

func main() {
	fmt.Println(numDecodings("12"))
}
