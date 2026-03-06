package containerwithmostwater

// [1,8,6,2,5,4,8,3,7]
// how to know which will contain most water
// for me first look its the highest and 2nd highest that will contain most, with the height
// of the second highest
// so we need to calculate the area and compare
// area = (index 1 - index 2) * min()
func maxArea(height []int) int {
	// condition for moving is increase which one is lower
	i, j := 0, len(height)-1
	var amount int
	for i < j {
		currentAmount := (j - i) * min(height[j], height[i])
		amount = max(amount, currentAmount)
		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}
	return amount
}
