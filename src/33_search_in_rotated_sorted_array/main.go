package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// 4 5 6 7 0 1 2
	// nums[left] = 0
	if nums[len(nums)-1] > target {
		right = left - 1
		left = 0
	} else {
		right = len(nums) - 1
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
