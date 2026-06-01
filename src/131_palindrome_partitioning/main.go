package main

import "fmt"

// "aab"
func partition(s string) [][]string {
	result := [][]string{}

	var recursivePartition func(index int, par []string)

	recursivePartition = func(index int, par []string) {
		if index == len(s) {
			result = append(result, append([]string{}, par...))
			return
		}

		for end := index + 1; end <= len(s); end++ {
			sub := s[index:end]
			if isPalindrom(sub) {
				recursivePartition(end, append(par, sub))
			}
		}
	}

	recursivePartition(0, []string{})

	return result
}

func isPalindrom(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
