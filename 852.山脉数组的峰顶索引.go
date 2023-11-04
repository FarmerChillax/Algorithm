/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 *
 * https://leetcode.cn/problems/peak-index-in-a-mountain-array/description/
 *
 * algorithms
 * Medium (68.51%)
 * Likes:    378
 * Dislikes: 0
 * Total Accepted:    146K
 * Total Submissions: 213.2K
 * Testcase Example:  '[0,1,0]'
 *
 * 符合下列属性的数组 arr 称为 山脉数组 ：
 *
 * arr.length >= 3
 * 存在 i（0 < i < arr.length - 1）使得：
 *
 * arr[0] < arr[1] < ... arr[i-1] < arr[i]
 * arr[i] > arr[i+1] > ... > arr[arr.length - 1]
 *
 *
 *
 *
 * 给你由整数组成的山脉数组 arr ，返回满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i +
 * 1] > ... > arr[arr.length - 1] 的下标 i 。
 *
 * 你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [0,1,0]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [0,2,1,0]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：arr = [0,10,5,2]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= arr.length <= 10^5
 * 0 <= arr[i] <= 10^6
 * 题目数据保证 arr 是一个山脉数组
 *
 *
 */
package algorithm

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	var left, right int = 0, len(arr) - 1
	for left < right {
		mid := (right + left) / 2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			// 峰顶
			return mid
		}
		if arr[mid] > arr[mid-1] && arr[mid+1] > arr[mid] {
			// 递增趋势，因此山峰在右边
			left = mid + 1
		} else {
			// 递减趋势，因此山峰在左侧
			// 山谷则默认往左走
			right = mid
		}
	}
	return -1
}

// @lc code=end
