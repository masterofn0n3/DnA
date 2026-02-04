package validanagram

func isAnagram(s string, t string) bool {
	//APPROACH1
	if len(s) != len(t) {
		return false
	}
	// charmap := map[rune]int{}
	// for _, v := range s {
	// 	charmap[v] += 1
	// }
	// for _, v := range t {
	// 	charmap[v] -= 1
	// }

	// for _, v := range charmap {
	// 	if v != 0 {
	// 		return false
	// 	}
	// }

	//26 english letter
	char := [26]int{}

	for i := range s {
		char[s[i]-'a']++
		char[t[i]-'a']--
	}
	for _, v := range char {
		if v != 0 {
			return false
		}
	}

	return true
}
