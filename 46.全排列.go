/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode.cn/problems/permutations/description/
 *
 * algorithms
 * Medium (78.95%)
 * Likes:    2732
 * Dislikes: 0
 * Total Accepted:    951K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[[0,1],[1,0]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * nums 中的所有整数 互不相同
 *
 *
 */
package algorithm

// @lc code=start
var (
	permuteRes  [][]int
	permutePath []int
	permuteUsed []bool
)

func permute(nums []int) [][]int {
	permuteRes = make([][]int, 0)
	permuteUsed = make([]bool, len(nums))
	backTracePermute(nums)
	return permuteRes
}

func backTracePermute(nums []int) {
	if len(permutePath) == len(nums) {
		permutePathCopy := make([]int, len(permutePath))
		copy(permutePathCopy, permutePath)
		permuteRes = append(permuteRes, permutePathCopy)
		return
	}
	for i := 0; i < len(nums); i++ {
		if permuteUsed[i] {
			continue
		}
		permutePath = append(permutePath, nums[i])
		permuteUsed[i] = true
		backTracePermute(nums)
		permuteUsed[i] = false
		permutePath = permutePath[:len(permutePath)-1]
	}
}

// @lc code=end
