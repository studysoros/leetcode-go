func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)

	for _, s := range strs {
		temp := [26]int{}
		for i := 0; i < len(s); i++ {
			temp[s[i]-'a'] += 1
		}
		mp[temp] = append(mp[temp], s)
	}
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}