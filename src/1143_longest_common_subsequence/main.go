package main

func longestCommonSubsequence(text1 string, text2 string) int {
	arr := make([]int, len(text2)+1)

	for i := len(text1) - 1; i >= 0; i-- {
		prev := 0
		for j := len(text2) - 1; j >= 0; j-- {
			temp := arr[j]
			if text1[i] == text2[j] {
				arr[j] = 1 + prev
			} else {
				arr[j] = max(arr[j], arr[j+1])
			}
			prev = temp
		}
	}

	return arr[0]

	// dp := make([][]int, len(text1)+1)
	// for i := range dp {
	//     dp[i] = make([]int, len(text2)+1)
	// }

	// for i := len(text1)-1; i >= 0; i-- {
	//     for j := len(text2)-1; j >= 0; j-- {
	//         if text1[i] == text2[j] {
	//             dp[i][j] = 1 + dp[i+1][j+1]
	//         } else {
	//             dp[i][j] = max(dp[i][j+1], dp[i+1][j])
	//         }
	//     }
	// }

	// return dp[0][0]

	// var dfs func(i, j int) int

	// dfs = func(i, j int) int {
	//     if i == len(text1) || j == len(text2) {
	//         return 0
	//     }

	//     if text1[i] == text2[j] {
	//         return 1 + dfs(i+1, j+1)
	//     }

	//     return  max(dfs(i, j+1), dfs(i+1, j))
	// }

	// return dfs(0, 0)
}
