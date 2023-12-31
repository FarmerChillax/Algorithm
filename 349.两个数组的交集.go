/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode.cn/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (74.18%)
 * Likes:    848
 * Dislikes: 0
 * Total Accepted:    495.9K
 * Total Submissions: 668.5K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[9,4]
 * 解释：[4,9] 也是可通过的
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 1000
 *
 *
 */
package algorithm

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool, len(nums1))
	for _, num := range nums1 {
		numMap[num] = true
	}

	ans := []int{}
	for index, num := range nums2 {
		if numMap[num] {
			ans = append(ans, nums2[index])
		}
		numMap[num] = false
	}
	return ans
}

// @lc code=end
