func groupAnagrams(strs []string) [][]string {
	Map := map[string][]string{}
	for _, value := range strs {
		tmp := []byte(value)
		slices.sort(tmp) //i直接修改
		sliced := string(tmp)
		Map[sliced] = append(Map[sliced], value)
	}
	return slices.Collect(maps.Values(Map))
}