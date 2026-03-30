package main

import "slices"

// 1 7
// 4
// 4 7
// 5 7
// 6 7
// 6 7
// if i set left = mid, the loop will stuck at six se7en, haha, get it...:>
// also if time > h, also means that the speed is too slow, so current mid value will be discard, so we set left = mid + 1
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, slices.Max(piles)
	for left < right {
		mid := (left + right) / 2
		time := 0
		for _, v := range piles {
			time += (v + mid - 1) / mid
		}
		if time > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
