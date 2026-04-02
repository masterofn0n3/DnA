package main

import "math"

func main() {
}

/*
a = 1 3 8 9
b = 2 4 5 7 10

=> 1 2 3 4 5 7 8 9 10 => med = 5

1 3 2 4 5 8 9 7 10
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	mergeLength := len(nums1) + len(nums2) // 9

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	left, right := 0, len(nums1)
	for left <= right {
		took1 := (left + right) / 2 // took 2

		took2 := (mergeLength+1)/2 - took1 //took 3

		left1, left2 := math.MinInt, math.MinInt
		right1, right2 := math.MaxInt, math.MaxInt

		if took1 > 0 {
			left1 = nums1[took1-1]
		}
		if took2 > 0 {
			left2 = nums2[took2-1]
		}
		if took1 < len(nums1) {
			right1 = nums1[took1]
		}
		if took2 < len(nums2) {
			right2 = nums2[took2]
		}

		if left1 > right2 {
			right = took1 - 1
		} else if left2 > right1 {
			left = took1 + 1
		} else {
			if mergeLength%2 == 0 {
				return (float64(max(left1, left2)) + float64(min(right1, right2))) / 2
			} else {
				return float64(max(left1, left2))
			}
		}
	}
	return 0
}
