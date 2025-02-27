/*
 * @lc app=leetcode.cn id=916 lang=golang
 *
 * [916] 单词子集
 *
 * https://leetcode.cn/problems/word-subsets/description/
 *
 * algorithms
 * Medium (47.20%)
 * Likes:    105
 * Dislikes: 0
 * Total Accepted:    11.5K
 * Total Submissions: 24.1K
 * Testcase Example:  '["amazon","apple","facebook","google","leetcode"]\n["e","o"]'
 *
 * 给你两个字符串数组 words1 和 words2。
 *
 * 现在，如果 b 中的每个字母都出现在 a 中，包括重复出现的字母，那么称字符串 b 是字符串 a 的 子集 。
 *
 *
 * 例如，"wrr" 是 "warrior" 的子集，但不是 "world" 的子集。
 *
 *
 * 如果对 words2 中的每一个单词 b，b 都是 a 的子集，那么我们称 words1 中的单词 a 是 通用单词 。
 *
 * 以数组形式返回 words1 中所有的 通用 单词。你可以按 任意顺序 返回答案。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 =
 * ["e","o"]
 *
 * 输出：["facebook","google","leetcode"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：words1 = ["amazon","apple","facebook","google","leetcode"], words2 =
 * ["lc","eo"]
 *
 * 输出：["leetcode"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：words1 = ["acaac","cccbb","aacbb","caacc","bcbbb"], words2 =
 * ["c","cc","b"]
 *
 * 输出：["cccbb"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= words1.length, words2.length <= 10^4
 * 1 <= words1[i].length, words2[i].length <= 10
 * words1[i] 和 words2[i] 仅由小写英文字母组成
 * words1 中的所有字符串 互不相同
 *
 *
 */
package algorithm

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "o"},
			},
			want: []string{"facebook", "google", "leetcode"},
		},
		{
			name: "case-2",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"lo", "eo"},
			},
			want: []string{"google", "leetcode"},
		},
		{
			name: "case-3",
			args: args{
				words1: []string{"acaac", "cccbb", "aacbb", "caacc", "bcbbb"},
				words2: []string{"c", "cc", "b"},
			},
			want: []string{"cccbb"},
		},
		// * 输入：words1 = ["acaac","cccbb","aacbb","caacc","bcbbb"], words2 =
		// * ["c","cc","b"]
		// *
		// * 输出：["cccbb"]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordSubsets(tt.args.words1, tt.args.words2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func cnt_freq(s string) []int {
// 	char_freq := make([]int, 26)
// 	for _, c := range s {
// 		char_freq[c-'a']++
// 	}
// 	return char_freq
// }

// func wordSubsets(A []string, B []string) []string {
// 	B_char_freq := make([]int, 26)
// 	for _, s := range B {
// 		v := cnt_freq(s)
// 		for i := 0; i <= 25; i++ {
// 			B_char_freq[i] = max(B_char_freq[i], v[i])
// 		}
// 	}
// 	res := make([]string, 0)
// 	for _, a := range A {
// 		A_char_freq := cnt_freq(a)
// 		flag := true
// 		for i := 0; i < 26; i++ {
// 			if A_char_freq[i] < B_char_freq[i] {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag == true {
// 			res = append(res, a)
// 		}
// 	}
// 	return res
// }
