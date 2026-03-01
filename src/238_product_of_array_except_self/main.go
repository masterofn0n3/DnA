package main

// I need some how to multiply all other index except the current index
// required O(n) so only one loop allow
// need a loop, what I need to do at current index?
// can I use a map to store the current indexed value?

// how to accumulate values as the loop goes on to each index
/*

	   1 2 3 4
left   1 1 2 6
right  24 12 4 1

	at 1 => 2*3
	at 2 => 1*3
	at 3 => 1*2

*/
func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	left[0] = 1
	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := range left {
		left[i] = left[i] * right[i]
	}
	return left
}
