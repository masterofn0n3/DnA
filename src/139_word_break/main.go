package main

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1) //
	dp[0] = true

	dict := map[string]struct{}{}

	for _, v := range wordDict {
		dict[v] = struct{}{}
	}

	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, ok := dict[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}
