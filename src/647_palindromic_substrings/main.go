package main

// a => 1
// ab => 2
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += longestAtIndex(i, i, s)
		count += longestAtIndex(i, i+1, s)
	}

	return count
}

func longestAtIndex(left, right int, s string) int {
	counter := 0
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return counter
		}
		counter++
		left--
		right++
	}

	return counter
}
