package main

func characterReplacement(s string, k int) int {
	//create a frequency map
	freqs := make(map[byte]int)
	res, i, maxFreq := 0, 0, 0

	for j := 0; j < len(s); j++ {
		// store the frequency of s[j], inclusive
		freqs[s[j]]++
		maxFreq = max(freqs[s[j]], maxFreq)

		// while the window is too large to fix, reduce it
		for (j-i+1)-maxFreq > k {
			freqs[s[i]]--
			i++
		}

		if j-i+1 > res {
			res = j - i + 1
		}
	}

	return res
}
