package main

// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tFreq := make(map[byte]int)
	sFreq := make(map[byte]int)
	result := ""
	for i := range t {
		tFreq[t[i]]++
		sFreq[s[i]]++
	}
	i := 0
	for j := len(t); j < len(s); j++ {
		if !isValid(tFreq, sFreq) {
			sFreq[s[j]]++
			continue
		}
		//ABC
		//DABCD j 4  i 2 4-1 2-1
		for isValid(tFreq, sFreq) {
			sFreq[s[i]]--
			i++
		}
		currentBest := s[i-1 : j+1]
		if result == "" || len(currentBest) < len(result) {
			result = currentBest
		}

	}
	return result
}

func isValid(t, s map[byte]int) bool {
	for k := range t {
		if t[k] != s[k] {
			return false
		}
	}
	return true
}

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}
