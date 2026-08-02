package main

func numDistinct(s string, t string) int {
	if len(t) > len(s) {
		return 0
	}

	arr := make([]int, len(t)+1)

	arr[len(t)] = 1

	/*
		s = "ab" t = "a" => count = 1
		arr[1] = 1
	*/
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				arr[j] += arr[j+1]
			}
		}
	}

	return arr[0]

	// dp := make([][]int, len(s)+1)
	// for i := range dp {
	//     dp[i] = make([]int, len(t)+1)
	//     dp[i][len(t)] = 1
	// }

	// for i := len(s)-1; i >= 0; i-- {
	//     for j := len(t)-1; j >= 0; j-- {
	//         if s[i] != t[j] {
	//             dp[i][j] += dp[i+1][j]
	//         } else {
	//         dp[i][j] += dp[i+1][j] + dp[i+1][j+1]
	//         }
	//     }
	// }

	// return dp[0][0]

	// memo := map[[2]int]int{}

	// var dfs func(i, j int) int

	// dfs = func(i, j int) int{
	// 	if i == len(s) && j < len(t) {
	// 		return 0
	// 	}
	// 	if j == len(t) {
	//         memo[[2]int{i, j}] = 1
	// 		return 1
	// 	}

	//     if v, ok := memo[[2]int{i,j}]; ok {
	//         return v
	//     }

	//     count := 0
	// 	if s[i] != t[j] {
	// 		count += dfs(i+1, j)
	// 	} else {
	// 		count += dfs(i+1, j)
	// 		count +=dfs(i+1, j+1)
	// 	}

	//     memo[[2]int{i,j}] = count

	//     return count

	// }
	// return dfs(0, 0)

}
