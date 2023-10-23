/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 *
 * https://leetcode.cn/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (71.24%)
 * Likes:    767
 * Dislikes: 0
 * Total Accepted:    310.3K
 * Total Submissions: 435.6K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
 *
 *
 * 只使用数字1到9
 * 每个数字 最多使用一次
 *
 *
 * 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 * 解释:
 * 1 + 2 + 4 = 7
 * 没有其他符合的组合了。
 *
 * 示例 2:
 *
 *
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 * 解释:
 * 1 + 2 + 6 = 9
 * 1 + 3 + 5 = 9
 * 2 + 3 + 4 = 9
 * 没有其他符合的组合了。
 *
 * 示例 3:
 *
 *
 * 输入: k = 4, n = 1
 * 输出: []
 * 解释: 不存在有效的组合。
 * 在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= k <= 9
 * 1 <= n <= 60
 *
 *
 */
package algorithm

// @lc code=start
var (
	res3  [][]int
	path3 []int
)

func combinationSum3(k int, n int) [][]int {
	path3, res3 = make([]int, 0, k), make([][]int, 0)
	backTrace3(n, k, 0, 1)
	return res3
}

func backTrace3(targetSum, k, sum, startIndex int) {
	if sum > targetSum {
		return
	}
	if k == len(path3) {
		// 终止条件
		if targetSum == sum {
			// 目标集合
			tmp := make([]int, k)
			copy(tmp, path3)
			res3 = append(res3, tmp)
		}
		return
	}
	for i := startIndex; i <= 9; i++ {
		sum += i
		path3 = append(path3, i)
		backTrace3(targetSum, k, sum, i+1)
		// 撤销回溯
		sum -= i
		path3 = path3[:len(path3)-1]

	}
}

// @lc code=end
