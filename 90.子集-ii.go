/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode.cn/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (63.52%)
 * Likes:    1164
 * Dislikes: 0
 * Total Accepted:    330K
 * Total Submissions: 519.7K
 * Testcase Example:  '[1,2,2]'
 *
 * 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,2]
 * 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 *
 *
 *
 *
 */
package algorithm

import "sort"

// @lc code=start

var (
	subsetWithDupRes  [][]int
	subsetWithDupPath []int
	used              []bool
)

func subsetsWithDup(nums []int) [][]int {
	subsetWithDupRes = make([][]int, 0)
	subsetWithDupPath = make([]int, 0, len(nums))
	used = make([]bool, len(nums))
	sort.Ints(nums)
	subsetWithDupBackTrace(nums, 0)
	return subsetWithDupRes
}

func subsetWithDupBackTrace(nums []int, startIndex int) {
	subsetWithDupPathCopy := make([]int, len(subsetWithDupPath))
	copy(subsetWithDupPathCopy, subsetWithDupPath)
	subsetWithDupRes = append(subsetWithDupRes, subsetWithDupPathCopy)

	if startIndex == len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		subsetWithDupPath = append(subsetWithDupPath, nums[i])
		used[i] = true
		subsetWithDupBackTrace(nums, i+1)
		// 撤销回溯
		used[i] = false
		subsetWithDupPath = subsetWithDupPath[:len(subsetWithDupPath)-1]
	}
}

// @lc code=end
