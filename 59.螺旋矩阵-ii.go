/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode.cn/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (71.91%)
 * Likes:    1184
 * Dislikes: 0
 * Total Accepted:    360.2K
 * Total Submissions: 500.9K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */
package algorithm

// @lc code=start
func generateMatrix(n int) [][]int {
	// 构建二维数组
	ans := make([][]int, n)
	for index := range ans {
		ans[index] = make([]int, n)
	}
	var offset int
	var num int
	var loopCnt int = n / 2
	var startX, startY int
	for loopCnt > 0 {
		for y := startY; y < n-1-offset; y++ {
			num++
			ans[startX][y] = num
		}

		for x := startX; x < n-1-offset; x++ {
			num++
			ans[x][n-1-offset] = num
		}

		for y := n - 1 - offset; y > offset; y-- {
			num++
			ans[n-1-startX][y] = num
		}

		for x := n - 1 - offset; x > offset; x-- {
			num++
			ans[x][startX] = num
		}

		// break
		startX++
		startY++
		offset++
		loopCnt--
	}

	if n%2 != 0 {
		mid := n / 2
		num++
		ans[mid][mid] = num
	}

	// for _, item := range ans {
	// 	fmt.Println(item)
	// }

	return ans
}

// @lc code=end
