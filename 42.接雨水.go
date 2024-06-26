/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode.cn/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (63.08%)
 * Likes:    4724
 * Dislikes: 0
 * Total Accepted:    772.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 *
 *
 * 示例 2：
 *
 *
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 1 <= n <= 2 * 10^4
 * 0 <= height[i] <= 10^5
 *
 *
 */

package algorithm

// @lc code=start
func trap(height []int) int {
	var leftHeight int
	for i := 0; i < len(height); i++ {
		if height[i] > 0 {
			leftHeight = i
			break
		}
		if i == len(height)-1 {
			// 整个数组都为 0
			return 0
		}
	}

	var ans int
	for i := 1; i < len(height); i++ {
		if height[i] >= height[leftHeight] {
			// 移动左侧指针
			leftHeight = i
		} else {
			ans += height[leftHeight] - height[i]
		}
	}
	return ans
}

// @lc code=end
