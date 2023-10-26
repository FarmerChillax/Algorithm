/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode.cn/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (73.39%)
 * Likes:    1663
 * Dislikes: 0
 * Total Accepted:    333.2K
 * Total Submissions: 454K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 *
 * 回文串 是正着读和反着读都一样的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a"
 * 输出：[["a"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由小写英文字母组成
 *
 *
 */
package algorithm

// @lc code=start
var (
	partitionRes  [][]string
	partitionPath []string
)

func partition(s string) [][]string {
	partitionRes = make([][]string, 0)
	partitionPath = make([]string, 0, len(s))

	backTracePartition(s, 0)
	return partitionRes
}

func backTracePartition(s string, startIndex int) {
	if startIndex >= len(s) {
		// 遍历完一次，找到了一组集合
		partitionPathCopy := make([]string, len(partitionPath))
		copy(partitionPathCopy, partitionPath)
		partitionRes = append(partitionRes, partitionPathCopy)
		return
	}

	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s[startIndex : i+1]) {
			partitionPath = append(partitionPath, s[startIndex:i+1])
			backTracePartition(s, i+1)
			// 撤销回溯
			partitionPath = partitionPath[:len(partitionPath)-1]
		}
	}

}

func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

// @lc code=end
