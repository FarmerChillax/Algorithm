/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode.cn/problems/non-decreasing-subsequences/description/
 *
 * algorithms
 * Medium (51.88%)
 * Likes:    737
 * Dislikes: 0
 * Total Accepted:    155.5K
 * Total Submissions: 300K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
 *
 * 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,6,7,7]
 * 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,4,3,2,1]
 * 输出：[[4,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 15
 * -100 <= nums[i] <= 100
 *
 *
 */
package algorithm

// @lc code=start

var (
	subsequenceRes  [][]int
	subsequencePath []int
)

func findSubsequences(nums []int) [][]int {
	subsequenceRes = make([][]int, 0)
	subsequencePath = make([]int, 0, len(nums))
	backTraceSubsequence(nums, 0)
	return subsequenceRes
}

func backTraceSubsequence(nums []int, startIndex int) {
	if len(subsequencePath) >= 2 {
		subsequencePathCopy := make([]int, len(subsequencePath))
		copy(subsequencePathCopy, subsequencePath)
		subsequenceRes = append(subsequenceRes, subsequencePathCopy)
	}

	if len(nums) == startIndex {
		return
	}
	usedSet := make(map[int]bool, len(nums))
	for i := startIndex; i < len(nums); i++ {
		if (len(subsequencePath) > 0 && nums[i] < subsequencePath[len(subsequencePath)-1]) ||
			usedSet[nums[i]] {
			continue
		}
		// 记录已经使用的元素
		usedSet[nums[i]] = true
		// 回溯前置
		subsequencePath = append(subsequencePath, nums[i])
		backTraceSubsequence(nums, i+1)
		subsequencePath = subsequencePath[:len(subsequencePath)-1]
	}
}

// @lc code=end
