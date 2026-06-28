package main

func maxProduct(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	maxProd, minProd := 1, 1
	result := 0

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		oldMax := maxProd
		maxProd = max(maxProd*current, current, minProd*current)
		minProd = min(oldMax*current, current, minProd*current)

		result = max(maxProd, result)
	}

	return result
}
