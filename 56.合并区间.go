/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode.cn/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (50.10%)
 * Likes:    2516
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 2.1M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]
 * 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 * 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[1,4],[4,5]]
 * 输出：[[1,5]]
 * 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= intervals.length <= 10^4
 * intervals[i].length == 2
 * 0 <= starti <= endi <= 10^4
 *
 *
 */
package algorithm

import "sort"

// @lc code=start
func merge(inputs [][]int) [][]int {
	result := [][]int{}
	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i][0] < inputs[j][0]
	})

	var left, right int = inputs[0][0], inputs[0][1]
	for i := 1; i < len(inputs); i++ {
		currLeft := inputs[i][0]
		currRight := inputs[i][1]

		// 如果左边的值比上一个区间的右边界要小，则合并
		if currLeft <= right {
			right = max(currRight, right)
		} else {
			// 否则表示区间断开，则存储结果
			result = append(result, []int{left, right})
			left = currLeft
			right = currRight
		}
	}
	result = append(result, []int{left, right})

	return result
}

// @lc code=end
