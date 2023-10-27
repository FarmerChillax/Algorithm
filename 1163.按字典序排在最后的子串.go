/*
 * @lc app=leetcode.cn id=1163 lang=golang
 *
 * [1163] 按字典序排在最后的子串
 *
 * https://leetcode.cn/problems/last-substring-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (35.08%)
 * Likes:    183
 * Dislikes: 0
 * Total Accepted:    22.6K
 * Total Submissions: 64.4K
 * Testcase Example:  '"abab"'
 *
 * 给你一个字符串 s ，找出它的所有子串并按字典序排列，返回排在最后的那个子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abab"
 * 输出："bab"
 * 解释：我们可以找出 7 个子串 ["a", "ab", "aba", "abab", "b", "ba", "bab"]。按字典序排在最后的子串是
 * "bab"。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "leetcode"
 * 输出："tcode"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 4 * 10^5
 * s 仅含有小写英文字符。
 *
 *
 */
package algorithm

import (
	"sort"
)

// @lc code=start
func lastSubstring(s string) string {
	sSet := make(map[string]bool)

	for i := 0; i <= len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			sSet[s[i:j]] = true
		}
	}
	res := make([]string, 0, len(sSet))
	for k := range sSet {
		res = append(res, k)
	}

	sort.Strings(res)

	return res[len(res)-1]
}

// @lc code=end
