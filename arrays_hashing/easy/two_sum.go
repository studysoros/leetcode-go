func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, v := range nums {
		if j, inHashMap := hashMap[target-v]; inHashMap {
			return []int{i, j}
		}
		hashMap[v] = i
	}

	return nil
}