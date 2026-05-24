package main

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets(nums []int) [][]int {
	result := [][]int{}
	var createSub func(index int, sub []int)

	createSub = func(index int, sub []int) {
		if index == len(nums) {
			//remember to create fresh copy here, or go memory behavior will mess with the result
			result = append(result, append([]int{}, sub...))
			return
		}
		includeSub := append(sub, nums[index])

		createSub(index+1, includeSub)
		createSub(index+1, sub)

	}

	createSub(0, []int{})

	return result
}
