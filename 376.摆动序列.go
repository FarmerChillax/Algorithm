/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 *
 * https://leetcode.cn/problems/wiggle-subsequence/description/
 *
 * algorithms
 * Medium (46.76%)
 * Likes:    1044
 * Dislikes: 0
 * Total Accepted:    205.5K
 * Total Submissions: 440K
 * Testcase Example:  '[1,7,4,9,2,5]'
 *
 * 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列
 * 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
 *
 *
 *
 * 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
 *
 * 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5]
 * 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
 *
 *
 * 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
 *
 * 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,7,4,9,2,5]
 * 输出：6
 * 解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,17,5,10,13,15,10,5,16,8]
 * 输出：7
 * 解释：这个序列包含几个长度为 7 摆动序列。
 * 其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,2,3,4,5,6,7,8,9]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度完成此题?
 *
 */

package algorithm

// @lc code=start
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			// （当前 diff 大于0 且 prev diff 小于 0）或
			// (当前 diff 小于0 且 prev diff 大于 0)
			ans++
			prevDiff = diff
		}
	}
	return ans
}

// @lc code=end
