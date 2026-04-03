package main

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
/*
	how to know if a string a not duplicated
	need a map to store frequency => if map[char] > 0 => duplicated
	advance j => added to the map
	a 1
	b 1
	c 1
	a already 1 => duplicated => but i don't know which position the last a belong to, do I need store it
	abca
	01

*/
func lengthOfLongestSubstring(s string) int {
	i, j := 0, 0
	longest := 0
	charMap := make(map[byte]struct{})
	for j < len(s) {
		if _, ok := charMap[s[j]]; !ok {
			charMap[s[j]] = struct{}{}
			j++
			currLength := j - i
			longest = max(currLength, longest)
		} else {
			delete(charMap, s[i])
			i++
		}
	}
	return longest
}
