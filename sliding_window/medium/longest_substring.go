func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	lastIndex := make([]int, 128)

	left := 0
	for right := 0; right < len(s); right++ {
		currentChar := s[right]
		if lastIndex[currentChar] > left {
			left = lastIndex[currentChar]
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		lastIndex[currentChar] = right + 1
	}

	return maxLen
}