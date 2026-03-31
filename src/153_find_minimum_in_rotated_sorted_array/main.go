package main

// 4,5,6,7,0,1,2
// 3 => 7
// 7> 2 => left = 3+1=4
// mid =  5 => 1
// 1 < 2
// right = 4
// 7,0,1,2,3,4,5,6
// 5,6,7,0,1,2,3,
// mid = 0

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
