package main

import "sort"

// Input: nums = [0,1,0,3,2,3]
// Output: 4
func lengthOfLIS(nums []int) int {
	tail := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i] > tail[len(tail)-1] {
			tail = append(tail, nums[i])
		} else {
			idx := sort.SearchInts(tail, nums[i])
			tail[idx] = nums[i]
		}
	}

	return len(tail)

}

// func lengthOfLIS(nums []int) int {
// 	dp := []int{1}

// 	for i := 1; i < len(nums); i++ {
// 		maxJ := 0
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				maxJ = max(dp[j], maxJ)
// 			}
// 		}
// 		dp = append(dp, maxJ+1)
// 	}
// 	longest := 0
// 	for _, v := range dp {
// 		longest = max(v, longest)
// 	}

// 	return longest
// }
