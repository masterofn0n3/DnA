package main

/*
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Arr, s2Arr := [26]int{}, [26]int{}

	for i := range s1 {
		s1Arr[s1[i]-'a']++
		s2Arr[s2[i]-'a']++
	}

	//3 6

	for i := 0; i < len(s2)-len(s1); i++ {
		if isEqual(s1Arr, s2Arr) {
			return true
		}
		s2Arr[s2[i+len(s1)]-'a']++
		s2Arr[s2[i]-'a']--
	}
	return isEqual(s1Arr, s2Arr)
}

func isEqual(s1A, s2A [26]int) bool {
	for i := range s1A {
		if s1A[i] != s2A[i] {
			return false
		}
	}
	return true
}
