func productExceptSelf(nums []int) []int {
	n := len(nums)

	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = 1

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			result[i] *= nums[j]
		}
	}
	return result
}
