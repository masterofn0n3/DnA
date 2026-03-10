package main

// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]
//
//	func dailyTemperatures(temperatures []int) []int {
//		answer := []int{}
//		for i := range temperatures {
//			answer = append(answer, 0)
//			for j := i + 1; j < len(temperatures); j++ {
//				if temperatures[i] < temperatures[j] {
//					answer[i] = j - i
//					break
//				}
//			}
//		}
//		return answer
//	}
//
// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]
/*
// maxStack [2,3,]
//answer [1,1,0,0,5]
*/
func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	maxStack := []int{}
	for i := range temperatures {
		for len(maxStack) > 0 && temperatures[i] > temperatures[maxStack[len(maxStack)-1]] {
			answer[maxStack[len(maxStack)-1]] = i - maxStack[len(maxStack)-1]
			maxStack = maxStack[0 : len(maxStack)-1]
		}
		maxStack = append(maxStack, i)
	}
	return answer
}
