package main

func main() {

}

func containsDuplicate(nums []int) bool {
	dupMap := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := dupMap[v]; ok {
			return true
		} else {
			dupMap[v] = struct{}{}
		}
	}
	return false
}
