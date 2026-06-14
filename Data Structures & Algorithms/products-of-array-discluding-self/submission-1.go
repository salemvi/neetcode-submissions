func productExceptSelf(nums []int) []int {
	n := len(nums)

	output := make([]int, n)

	for i := 0; i < n; i++ {
		output[i] = 1

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			output[i] *= nums[j]
		}
	}
	return output
}
