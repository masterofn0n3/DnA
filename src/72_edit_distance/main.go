package main

func minDistance(word1 string, word2 string) int {

	arr := make([]int, len(word2)+1)
	// arr[i] represent the base case of when i == len(word1)
	for i := range arr {
		arr[i] = len(word2) - i
	}

	for i := len(word1) - 1; i >= 0; i-- {
		prev := arr[len(word2)]          //init to the lower right point of the matrix
		arr[len(word2)] = len(word1) - i // update to dp[i][n] for this row's own boundary

		for j := len(word2) - 1; j >= 0; j-- {
			temp := arr[j] // dp[i+1][j], old value before overwrite
			if word1[i] == word2[j] {
				arr[j] = prev
			} else {
				arr[j] = 1 + min(prev, temp, arr[j+1])
			}
			prev = temp // carry the OLD value forward as next diag, not the new one
		}
	}

	return arr[0]

	// dp := make([][]int, len(word1)+1)
	// for i := range dp {
	//     dp[i] = make([]int, len(word2)+1)
	//     dp[i][len(word2)] = len(word1) - i
	// }

	// for j := range dp[len(word1)]{
	//     dp[len(word1)][j] = len(word2) - j
	// }

	// for i := len(word1)-1; i >= 0; i-- {
	//     for j := len(word2)-1; j >= 0; j-- {
	//         if word1[i] == word2[j] {
	//             dp[i][j] = dp[i+1][j+1]
	//         } else {
	//             dp[i][j] = 1 + min(dp[i+1][j+1], dp[i][j+1], dp[i+1][j])
	//         }
	//     }
	// }

	// return dp[0][0]

	// var dfs func(i, j int) int

	// dfs = func(i, j int) int {
	//     if i == len(word1) {
	//         return len(word2) - j
	//     }
	//     if j == len(word2) {
	//         return len(word1) - i
	//     }

	//     if word1[i] == word2[j] {
	//         return dfs(i+1, j+1)
	//     }

	//     return 1+min(dfs(i+1, j+1), dfs(i, j+1), dfs(i+1, j))
	// }

	// return dfs(0, 0)
}
