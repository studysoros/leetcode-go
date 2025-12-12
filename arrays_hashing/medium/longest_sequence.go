func longestConsecutive(nums []int) int {
	numsSet := make(map[int]struct{})
	for _, v := range nums {
		numsSet[v] = struct{}{}
	}

	maxLen := 0

	for num := range numsSet {
		if _, ok := numsSet[num-1]; !ok {
			length := 1
			for {
				if _, ok := numsSet[num+length]; ok {
					length++
					continue
				}
				break
			}
			maxLen = max(maxLen, length)
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}