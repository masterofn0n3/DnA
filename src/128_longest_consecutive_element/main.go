package main

//100,4,200,1,3,2

func longestConsecutive(nums []int) int {
	longest := 0
	numMap := map[int]struct{}{}
	for _, v := range nums {
		numMap[v] = struct{}{}
	}
	for k := range numMap {
		if _, ok := numMap[k-1]; !ok {
			currentLen := 1
			for i := 1; i < len(nums); i++ {
				if _, ok := numMap[k+i]; ok {
					currentLen++
				} else {
					break
				}
			}
			if currentLen > longest {
				longest = currentLen
			}
		}
	}

	return longest
}
