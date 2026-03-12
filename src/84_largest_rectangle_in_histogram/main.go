package main

/*
Input: heights = [2,1,5,6,2,3]
next right smaller = [1,0,5,5,0,0]
stack [1]
Output: 10

I want to test out some new way of thinking
Pretent I'm the current element in the for loop
What do I know
I know my value is 2
I know the array length is 6
I respresent the height of a column in a histogram
A column will have width of 1 => so the area I could make is my height * 1 = 2
A larger area could form by extending to the adjacent element for example the one next to me, 1
But the height will be limit by the lowest column in the group, so 1 * outer right index - outer left index, in this case: 2
so 2 * 1 = 1
To keep track of largest area I'm forming, I need a value, so `area`

What do I want to know
I want to know which was the largest area I can form with other element
*/
func largestRectangleArea(heights []int) int {
	rightStack := []int{}
	nextRightSmaller := make([]int, len(heights))
	for i := range heights {
		for len(rightStack) > 0 && heights[i] < heights[rightStack[len(rightStack)-1]] {
			nextRightSmaller[rightStack[len(rightStack)-1]] = i
			rightStack = rightStack[:len(rightStack)-1]
		}
		rightStack = append(rightStack, i)
	}
	for _, v := range rightStack {
		nextRightSmaller[v] = len(heights)
	}
	leftStack := []int{}
	nextLeftSmaller := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for len(leftStack) > 0 && heights[i] < heights[leftStack[len(leftStack)-1]] {
			nextLeftSmaller[leftStack[len(leftStack)-1]] = i
			leftStack = leftStack[:len(leftStack)-1]
		}
		leftStack = append(leftStack, i)
	}
	for _, v := range leftStack {
		nextLeftSmaller[v] = -1
	}
	maxArea := 0
	for i, v := range heights {
		area := (nextRightSmaller[i] - nextLeftSmaller[i] - 1) * v
		maxArea = max(area, maxArea)
	}
	return maxArea
}
