package threesome

import (
	"slices"
)

// things to solve
// sort
// how to recognize duplicated triplet
// -1,0,1,2,-1,-4
// -4 -1 -1 0 1 2
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}
	for i := 0; i <= len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[j]+nums[k] < -nums[i] {
				j++
			} else if nums[j]+nums[k] > -nums[i] {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			}
		}
	}
	return result
}
