package main

func climbStairs(n int) int {
	secondLast, last := 0, 1

	for i := 1; i <= n; i++ {
		secondLast, last = last, secondLast+last
	}

	return last

}

// func climbStairs(n int) int {
// 	res, prev, nextPrev := 0, 0, 0

// 	for i := 1; i <= n; i++ {
// 		if i == 1 {
// 			res = 1
// 		} else if i == 2 {
// 			res = 2
// 		} else {
// 			res = prev + nextPrev
// 		}

// 		nextPrev = prev
// 		prev = res

// 	}

// 	return res

// }

// func climbStairs(n int) int {
// 	wayMap := map[int]int{}
// 	var climb func(sum int) int
// 	climb = func(sum int) int {
// 		if sum == n {
// 			return 1
// 		}

// 		if sum > n {
// 			return 0
// 		}

// 		if _, ok := wayMap[sum]; !ok {
// 			way1 := climb(sum + 1)
// 			way2 := climb(sum + 2)

// 			wayMap[sum] = way1 + way2
// 		}

// 		return wayMap[sum]
// 	}

// 	return climb(0)

// }
