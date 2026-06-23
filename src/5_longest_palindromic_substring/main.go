package main

// i = 1, longest is 1 a
// i = 2, longest is 2 b
// i = 3 e.g abc => longest is 2, if aba then longest is 3
func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		start, end := longestAtIndex(i, i, s)
		if end-start+1 > len(result) {
			result = s[start : end+1]
		}
		start, end = longestAtIndex(i, i+1, s)
		if end >= start && end-start+1 > len(result) {
			result = s[start : end+1]
		}
	}

	return result
}

func longestAtIndex(left, right int, s string) (int, int) {

	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return left + 1, right - 1
		}
		left--
		right++
	}

	return left + 1, right - 1
}
