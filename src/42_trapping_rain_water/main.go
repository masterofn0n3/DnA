package main

/*
[0,1,0,2]
[0,0,1,2]

how to know how much water it can trap
water can be trap between 2 index if there's a low point between them
but the height varies, so water can be trap at different places
*/
// func trap(height []int) int {
// 	totalWater := 0
// 	for i := 1; i < len(height)-1; i++ {
// 		maxLeft, maxRight := 0, 0
// 		for j := i - 1; j > 0; j-- {
// 			maxLeft = max(height[j], maxLeft)
// 		}
// 		for k := i + 1; k < len(height); k++ {
// 			maxRight = max(height[k], maxRight)
// 		}
// 		waterLevel := min(maxLeft, maxRight)
// 		if waterLevel > height[i] {
// 			totalWater += waterLevel - height[i]
// 		}
// 	}
// 	return totalWater
// }

// [0,1,0,2]
// [0,0,1,2]
// [  2 ,2,2,0]
// is this different from your propose approach

// No, it’s the same idea—just a different (and actually nicer) convention.
// What I described:
// leftMax[i] = max of height[0..i] (includes i).
// “Max to the left of i” = leftMax[i-1].
// What you did:
// leftMax[i] = max of height[0..i-1] (only to the left of i).
// So “max to the left of i” is exactly leftMax[i].
// So you’re storing “max to the left of i” in leftMax[i] directly, and using leftMax[i] in the water step. I had “max up to and including i” in leftMax[i] and used leftMax[i-1] for the water. The water formula is the same; only the meaning of the index differs.
// Your version is consistent and avoids the “use i-1 when reading” detail. So: same approach, same complexity, your indexing is just a cleaner choice for this problem.
func trap(height []int) int {
	leftMax := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i-1], leftMax[i-1])
	}
	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = 0
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i+1])
	}
	totalWater := 0
	for i := range height {
		waterLevel := min(leftMax[i], rightMax[i])
		if waterLevel > height[i] {
			totalWater += waterLevel - height[i]
		}
	}
	return totalWater
}
