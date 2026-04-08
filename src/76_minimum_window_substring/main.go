package main

// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}
	result := ""
	uniqueChar := len(tMap)
	i := 0
	count := 0
	for j := 0; j < len(s); j++ {
		sMap[s[j]]++
		if v, ok := tMap[s[j]]; ok && v == sMap[s[j]] {
			count++
		}
		if count < uniqueChar {
			continue
		}

		for count == uniqueChar {
			best := s[i : j+1]
			if result == "" || len(best) < len(result) {
				result = best
			}
			sMap[s[i]]--
			if v, ok := tMap[s[i]]; ok && v > sMap[s[i]] {
				count--
			}
			i++
		}

	}
	return result
}

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}
