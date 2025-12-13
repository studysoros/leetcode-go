func characterReplacement(s string, k int) int {
	freqs := make(map[byte]int)
	var res, left, maxFreq int

	for right := 0; right < len(s); right++ {
		freqs[s[right]]++
		if freqs[s[right]] > maxFreq {
			maxFreq = freqs[s[right]]
		}

		for (right-left+1)-maxFreq > k {
			freqs[s[left]]--
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}

	return res
}

// Solution 2
func characterReplacement(s string, k int) int {
	freqs := make(map[byte]int)
	var left, right, res, maxFreq int

	for right < len(s) {
		freqs[s[right]]++
		if freqs[s[right]] > maxFreq {
			maxFreq = freqs[s[right]]
		}

		if (right-left+1)-maxFreq > k {
			freqs[s[left]]--
			left++
		}

		if (right - left + 1) > res {
			res = right - left + 1
		}

		right++
	}

	return res
}