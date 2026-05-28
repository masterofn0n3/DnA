package main

import "sort"

// Example 1:

// Input: candidates = [10,1,2,7,6,1,5], target = 8
// Output:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	// 1 1 2 5 6 7 10
	result := [][]int{}

	var comSum func(index, currentSum int, sumArr []int)

	comSum = func(index, currentSum int, sumArr []int) {
		if currentSum == target {
			result = append(result, append([]int{}, sumArr...))
			return
		} else if currentSum > target {
			return
		}
		if index > len(candidates)-1 {
			return
		}

		comSum(index+1, currentSum+candidates[index], append(sumArr, candidates[index]))
		// 1 => 1 1 => 1 1 2
		nextIdx := index + 1
		for nextIdx < len(candidates) && candidates[nextIdx] == candidates[index] {
			nextIdx++
		}
		comSum(nextIdx, currentSum, sumArr)

	}

	comSum(0, 0, []int{})

	return result
}
