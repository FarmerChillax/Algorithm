/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (58.49%)
 * Likes:    2659
 * Dislikes: 0
 * Total Accepted:    766.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */

package algorithm

// @lc code=start

var digitsMap []string = []string{
	"", "", // index: 0, 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

var (
	letterRes  []string
	letterPath string
)

func letterCombinations(digits string) []string {
	letterRes = make([]string, 0)
	letterPath = ""
	if len(digits) == 0 {
		return letterRes
	}

	// 转换为数组
	matrix := make([]string, 0, len(digits))
	for _, digit := range digits {
		btn := int(digit - '0')
		matrix = append(matrix, digitsMap[btn])
	}
	// 开始回溯
	backTraceLetterCombination(matrix, 0)
	return letterRes
}

func backTraceLetterCombination(digitMatrix []string, startIndex int) {
	if len(digitMatrix) == len(letterPath) {
		letterRes = append(letterRes, letterPath)
		return
	}

	for _, digit := range digitMatrix[startIndex] {
		letterPath += string(digit)
		backTraceLetterCombination(digitMatrix, startIndex+1)
		// 撤销回溯
		letterPath = letterPath[:len(letterPath)-1]
	}

}

// @lc code=end
