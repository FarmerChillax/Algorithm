/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode.cn/problems/combinations/description/
 *
 * algorithms
 * Medium (77.08%)
 * Likes:    1514
 * Dislikes: 0
 * Total Accepted:    606.9K
 * Total Submissions: 787.4K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
 *
 * 你可以按 任何顺序 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 4, k = 2
 * 输出：
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 * 示例 2：
 *
 *
 * 输入：n = 1, k = 1
 * 输出：[[1]]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */
package algorithm

// @lc code=start
var (
	res  [][]int
	path []int
)

func combine(n int, k int) [][]int {
	// 初始化全局变量
	path, res = make([]int, 0, k), make([][]int, 0)
	// 开始回溯
	backTrace(n, k, 1)
	return res
}

func backTrace(n, k, startIndex int) {
	if k == len(path) {
		// 创建新的数组，防止引用问题
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := startIndex; i <= n-(k-len(path))+1; i++ {
		// k - len(path) 为 <还需要的长度>
		// n - <还需要长度> -> 起始点
		// 如果 startIndex 超过了起始点
		// 则永远达不到目标长度，则提前返回（剪枝）
		path = append(path, i)
		backTrace(n, k, i+1)
		path = path[:len(path)-1]
	}
}

// @lc code=end
