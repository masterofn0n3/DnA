package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	var recurseSubset func(index int, sub []int)

	recurseSubset = func(index int, sub []int) {
		if index == len(nums) {
			result = append(result, append([]int{}, sub...))
			return
		}

		recurseSubset(index+1, append(sub, nums[index]))
		nextIndex := index + 1
		for nextIndex < len(nums) && nums[nextIndex] == nums[index] {
			nextIndex++
		}
		recurseSubset(nextIndex, sub)

	}

	recurseSubset(0, []int{})

	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{5, 5, 5, 5, 5}))

}
