package sumii

// num 1 + num 2 = target
// num 1 = target - num2
// 2,7,11,15 target 9
func twoSumFirstApproach(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[j]+numbers[i] > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}
