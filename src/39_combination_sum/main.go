package main

// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.
func combinationSum(candidates []int, target int) [][]int {
	//stop when the sum is equal to the target
	//the decision is whether to add the current number or not
	//but we can also add a number many times?, how to include this decision
	result := [][]int{}

	var sums func(index int, sumArr []int, currentSum int)

	sums = func(index int, sumArr []int, currentSum int) {
		if currentSum == target {
			result = append(result, append([]int{}, sumArr...))
			return
		} else if currentSum > target {
			return
		}
		if index > len(candidates)-1 {
			return
		}
		sums(index, append(sumArr, candidates[index]), currentSum+candidates[index]) // 2, [2]
		// sums(index+1, append(sumArr, candidates[index]))
		sums(index+1, sumArr, currentSum) //3 []

	}

	sums(0, []int{}, 0)

	return result
}
