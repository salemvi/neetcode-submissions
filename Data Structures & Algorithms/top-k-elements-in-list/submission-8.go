func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	bucket := make([][]int, n + 1)

	for num, count := range countMap {
		bucket[count] = append(bucket[count], num)
	}

	res := make([]int, 0, k)

	for i := n; i > 0 && len(res) < k; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)

			if len(res) == k {
				return res
			}
		}
	}
	return res
}
