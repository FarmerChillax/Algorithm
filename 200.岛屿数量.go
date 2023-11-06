/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode.cn/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (59.68%)
 * Likes:    2342
 * Dislikes: 0
 * Total Accepted:    725.3K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 *
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 *
 * 此外，你可以假设该网格的四条边均被水包围。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1
 * grid[i][j] 的值为 '0' 或 '1'
 *
 *
 */
package algorithm

// @lc code=start
func numIslands(grid [][]byte) int {
	var ans int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '1' {
				// 找到岛屿
				ans++
				// 淹没岛屿
				dfsNumIslands(grid, x, y)
			}
		}
	}
	return ans
}

func dfsNumIslands(grid [][]byte, x, y int) {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
		// 下标越界
		return
	}
	if grid[y][x] == '0' {
		// 如果已经是海水了，则返回
		return
	}
	// 将当前岛屿淹没
	grid[y][x] = '0'
	dfsNumIslands(grid, x, y+1)
	dfsNumIslands(grid, x, y-1)
	dfsNumIslands(grid, x+1, y)
	dfsNumIslands(grid, x-1, y)
}

// @lc code=end
