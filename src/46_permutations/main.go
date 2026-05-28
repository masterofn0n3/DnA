package main

// Input: nums = [1,2,3]
// Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	result := [][]int{}
	usageArr := make([]bool, len(nums))

	var recursivePerm func(perm []int)

	recursivePerm = func(perm []int) {
		if len(perm) == len(nums) {
			result = append(result, append([]int{}, perm...))
			return
		}
		for i, v := range usageArr {
			if !v {
				usageArr[i] = true
				perm = append(perm, nums[i])
				recursivePerm(perm)
				usageArr[i] = false
				perm = perm[:len(perm)-1]
			}
		}
	}

	recursivePerm([]int{})
	return result
}
