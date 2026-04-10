package main

// Example 1:

// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]
func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	deq := []int{}
	for i := 0; i < len(nums); i++ {
		for len(deq) > 0 && nums[i] > nums[deq[len(deq)-1]] {
			//to remove the last element use :len(arr)-1 not len(arr)
			deq = deq[:len(deq)-1]
		}
		//1
		//3

		deq = append(deq, i)

		if deq[0] == i-k {
			deq = deq[1:]
		}

		if i >= k-1 {
			result = append(result, nums[deq[0]])
		}
	}
	return result
}
