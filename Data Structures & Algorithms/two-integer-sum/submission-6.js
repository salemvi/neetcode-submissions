class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        const seen = new Map()

        for (let i = 0; i < nums.length; i++) {
            const x = nums[i]
            const need = target - x

            if (seen.has(need)) {
                const j = seen.get(need)
                return j < i ? [j, i]: [i, j]
            }

            seen.set(x, i)
        }
        return []
    }
}
