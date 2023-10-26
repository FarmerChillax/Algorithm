/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 *
 * https://leetcode.cn/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (58.53%)
 * Likes:    1333
 * Dislikes: 0
 * Total Accepted:    372.7K
 * Total Submissions: 636.5K
 * Testcase Example:  '"25525511135"'
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
 * 和 "192.168@1.1" 是 无效 IP 地址。
 *
 *
 * 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能
 * 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * s 仅由数字组成
 *
 *
 */
package algorithm

import (
	"strings"
)

// @lc code=start

var (
	ipAddressRes  []string
	ipAddressPath []string
)

func restoreIpAddresses(s string) []string {
	ipAddressRes = make([]string, 0)
	ipAddressPath = make([]string, 0)
	backTraceIpAddress(&s, 0, -1)
	return ipAddressRes
}

func backTraceIpAddress(s *string, startIndex, pointNum int) {
	// 字符串中有4段，则表示需要 return 了
	if len(ipAddressPath) == 4 {
		if startIndex == len(*s) {
			ipAddressRes = append(ipAddressRes, strings.Join(ipAddressPath, "."))
		}
		return
	}

	for i := startIndex; i < len(*s); i++ {
		if !isValid(s, startIndex, i) {
			// 当前子串不合法，则再加上后续的子串也会不合法
			// 因此终止查找
			return
		}
		// 满足子串则拼接
		ipAddressPath = append(ipAddressPath, (*s)[startIndex:i+1])
		backTraceIpAddress(s, i+1, pointNum+1)
		// 撤销回溯
		ipAddressPath = ipAddressPath[:len(ipAddressPath)-1]
	}

}

func isValid(s *string, start, end int) bool {
	if start > end {
		return false
	}

	if (*s)[start] == '0' && start != end {
		return false
	}

	var num int
	for i := start; i <= end; i++ {
		if (*s)[i] > '9' || (*s)[i] < '0' {
			return false
		}
		num = num*10 + int((*s)[i]-'0')
		if num > 255 {
			return false
		}
	}

	return true
}

// @lc code=end
