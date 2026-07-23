func hasDuplicate(nums []int) bool {
    if len(nums) < 2 {
        return false
    }

    seen := make(map[int]struct{})

    for _, num := range nums {
        if _, ok := seen[num]; ok {
            return true;
        }

        seen[num] = struct{}{}
    }
    return false
}
