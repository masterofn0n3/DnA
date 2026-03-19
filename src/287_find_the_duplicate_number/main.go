package main

/*
[1,3,4,2,2]
slow 1 fast 3
slow 3 fast 4
slow 2 fast 4
slow 4 fast 4
*/
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
