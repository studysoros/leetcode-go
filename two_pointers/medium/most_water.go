func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var mostArea int

	for left < right {
		currArea := min(height[left], height[right]) * (right - left)
		mostArea = max(mostArea, currArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return mostArea
}