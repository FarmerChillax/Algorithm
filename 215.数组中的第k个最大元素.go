package algorithm

import (
	"math"
)

// @lc code=start
func findKthLargest(nums []int, k int) int {
	kMap := make(map[int]bool, k)
	var max int = math.MinInt
	for k > 0 {
		max = math.MinInt
		var maxIdx int
		for idx, num := range nums {
			if num > max && !kMap[idx] {
				max = nums[idx]
				maxIdx = idx
			}
		}
		kMap[maxIdx] = true
		k--
	}
	return max
}

// @lc code=end
