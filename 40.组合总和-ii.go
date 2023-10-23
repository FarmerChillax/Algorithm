/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode.cn/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (59.55%)
 * Likes:    1470
 * Dislikes: 0
 * Total Accepted:    471.3K
 * Total Submissions: 791.7K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */
package algorithm

import "sort"

// @lc code=start
var (
	sumRes2  [][]int
	sumPath2 []int
)

func combinationSum2(candidates []int, target int) [][]int {
	sumPath2, sumRes2 = make([]int, 0), make([][]int, 0)
	sort.Ints(candidates) // 排序，为剪枝做准备

	backTraceSum2(candidates, target, 0)
	return sumRes2
}

func backTraceSum2(candidates []int, target, startIndex int) {
	// 回溯终止条件
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(sumPath2))
		copy(tmp, sumPath2)
		sumRes2 = append(sumRes2, tmp)
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if i != startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		sumPath2 = append(sumPath2, candidates[i])
		backTraceSum2(candidates, target-candidates[i], i+1)
		sumPath2 = sumPath2[:len(sumPath2)-1]
	}
}

// @lc code=end
