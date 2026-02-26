package groupanagram

import "sort"

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	// for pos := 0; pos < len(strs); pos++ {
	// 	subRes := []string{strs[pos]}
	// 	for j := pos + 1; j < len(strs); j++ {
	// 		//remove the added anagram
	// 		if isAnagram(strs[pos], strs[j]) {
	// 			subRes = append(subRes, strs[j])
	// 			strs = append(strs[:j], strs[j+1:]...)
	// 			j--
	// 		}
	// 	}
	// 	strs = append(strs[:pos], strs[pos+1:]...)
	// 	pos--
	// 	result = append(result, subRes)
	// }

	anagramMap := make(map[string][]string)
	for _, v := range strs {
		sortVal := sortString(v)
		anagramMap[sortVal] = append(anagramMap[sortVal], v)
	}
	for _, v := range anagramMap {
		result = append(result, v)
	}
	return result
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b) // "aet"
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	var charMap [26]int
	for i := range str1 {
		charMap[str1[i]-'a'] += 1
		charMap[str2[i]-'a'] -= 1
	}
	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}
	return true
}
