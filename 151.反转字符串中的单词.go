/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 *
 * https://leetcode.cn/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (52.95%)
 * Likes:    1045
 * Dislikes: 0
 * Total Accepted:    462.7K
 * Total Submissions: 873.2K
 * Testcase Example:  '"the sky is blue"'
 *
 * 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
 *
 * 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
 *
 * 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
 *
 * 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "the sky is blue"
 * 输出："blue is sky the"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "  hello world  "
 * 输出："world hello"
 * 解释：反转后的字符串中不能存在前导空格和尾随空格。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a good   example"
 * 输出："example good a"
 * 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 包含英文大小写字母、数字和空格 ' '
 * s 中 至少存在一个 单词
 *
 *
 *
 *
 *
 *
 *
 * 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
 *
 */
package algorithm

// @lc code=start
func reverseWords(s string) string {
	strBytes := []byte(s)

	// 移除多余空格
	var slow int
	for i := 0; i < len(strBytes); i++ {
		if strBytes[i] != ' ' {
			// 不是空格，则触发迁移规则
			if slow != 0 {
				strBytes[slow] = ' '
				slow++
			}
			for i < len(strBytes) && strBytes[i] != ' ' {
				strBytes[slow] = strBytes[i]
				slow++
				i++
			}
		}
	}
	// 截取有效长度
	strBytes = strBytes[:slow]
	// 反转字符串
	reverseS(strBytes, 0, len(strBytes))
	// 反转单词
	slow = 0
	for i := 0; i < len(strBytes); i++ {
		if strBytes[i] == ' ' {
			reverseS(strBytes, slow, i)
			slow = i + 1
		}
	}

	reverseS(strBytes, slow, len(strBytes))

	return string(strBytes)

}

func reverseS(s []byte, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// @lc code=end
