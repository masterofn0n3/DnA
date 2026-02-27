package topkfrequentelement

// I can loop and use a map to store the frequency of the unique element
// but I still need to sort the map to know which have the most appearance
// is there a way to group and know the order in one loop?
func topKFrequent(nums []int, k int) []int {
	fMap := map[int]int{}
	for _, v := range nums {
		fMap[v]++
	}

	fArr := make([][]int, len(nums)+1)
	for k, v := range fMap {
		fArr[v] = append(fArr[v], k)
	}

	var result []int
	for i := len(fArr) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, fArr[i]...)
	}

	return result
}
