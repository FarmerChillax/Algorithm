/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (39.27%)
 * Likes:    9733
 * Dislikes: 0
 * Total Accepted:    2.6M
 * Total Submissions: 6.5M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */
package algorithm

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	charMap := make(map[byte]int, len(s))
	var ans, slow int
	for fast := 0; fast < len(s); fast++ {
		idx, ok := charMap[s[fast]]
		if ok && idx >= slow && idx != fast {
			// 如果记录的下标不一致，则表示子串中出现过该字符
			// 移动窗口边界到该字符后一位
			slow = charMap[s[fast]] + 1
		}
		charMap[s[fast]] = fast
		subLen := fast - slow + 1
		ans = max(ans, subLen)
	}

	return ans
}

// func max(a, b int) int {
// 	if b > a {
// 		return b
// 	}
// 	return a
// }

// @lc code=end
