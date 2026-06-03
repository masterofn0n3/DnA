package main

func letterCombinations(digits string) []string {
	result := []string{}

	digitArr := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var combine func(index int, com string)
	combine = func(index int, com string) {
		if index == len(digits) {
			result = append(result, com)
			return
		}
		for _, v := range digitArr[digits[index]] {
			combine(index+1, com+string(v))
		}
	}
	combine(0, "")

	return result
}
