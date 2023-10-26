/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (81.15%)
 * Likes:    2173
 * Dislikes: 0
 * Total Accepted:    690.6K
 * Total Submissions: 851K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
 * nums 中的所有元素 互不相同
 *
 *
 */
package algorithm

// @lc code=start
var (
	subsetRes  [][]int
	subsetPath []int
)

func subsets(nums []int) [][]int {
	subsetRes = make([][]int, 0)
	subsetPath = make([]int, 0, len(nums))
	subsetBackTrace(nums, 0)
	return subsetRes
}

func subsetBackTrace(nums []int, startIndex int) {
	subsetPathCopy := make([]int, len(subsetPath))
	copy(subsetPathCopy, subsetPath)
	subsetRes = append(subsetRes, subsetPathCopy)
	if startIndex == len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		subsetPath = append(subsetPath, nums[i])
		subsetBackTrace(nums, i+1)
		subsetPath = subsetPath[:len(subsetPath)-1]
	}
}

// @lc code=end
