func groupAnagrams(strs []string) [][]string {
	mm := make(map[string][]string)

	for _, w := range strs {
		key := sortString(w)
		mm[key] = append(mm[key], w)
	}

	res := make([][]string, 0, len(mm))

	for _, group := range mm {
		res = append(res, group)
	}

	return res

}

func sortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})

	return string(runes)
}