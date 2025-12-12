func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for key, v := range mp {
		bucket[v] = append(bucket[v], key)
	}

	ans := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, v := range bucket[i] {
			if k > 0 {
				ans = append(ans, v)
				k--
			}
		}
	}
	return ans
}

//  Solution 2
func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for key, v := range mp {
		bucket[v] = append(bucket[v], key)
	}

	ans := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, v := range bucket[i] {
			ans = append(ans, v)
			if len(ans) == k {
				return ans
			}
		}
	}

	return nil
}