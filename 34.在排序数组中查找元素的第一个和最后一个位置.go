/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.58%)
 * Likes:    2505
 * Dislikes: 0
 * Total Accepted:    862.8K
 * Total Submissions: 2M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
 *
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 *
 * 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 *
 * 示例 3：
 *
 *
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums 是一个非递减数组
 * -10^9 <= target <= 10^9
 *
 *
 */
package algorithm

// @lc code=start
func searchRange(nums []int, target int) []int {
	var left, right int = 0, len(nums) - 1
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] == target {
			// 找到目标，开始查找边界
			left, right = mid, mid
			for left >= 0 && nums[left] == target {
				left--
			}
			for right <= len(nums)-1 && nums[right] == target {
				right++
			}
			return []int{left + 1, right - 1}
		}
		if nums[mid] > target {
			// 当前元素大于目标元素，因此目标元素在左边 [left, mid-1]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return []int{-1, -1}
}

func getLeftBorder(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	leftBorder := -2
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] >= target {
			// 当前元素大于目标元素，因此目标元素在左边 [left, mid-1]
			// 获取左边界
			right = mid - 1
			if nums[mid] == target {
				leftBorder = right
			}
		} else {
			left = mid + 1
		}
	}
	return leftBorder
}

func getRightBorder(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	border := -2
	for left <= right {
		mid := left + ((right - left) / 2)
		if nums[mid] <= target {
			// 当前元素大于目标元素，因此目标元素在左边 [left, mid-1]
			// 获取左边界
			left = mid + 1
			if nums[mid] == target {
				border = left
			}
		} else {
			right = mid - 1
		}
	}
	return border
}

// @lc code=end
