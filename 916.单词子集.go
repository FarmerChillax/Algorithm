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

// @lc code=start
func wordSubsets(words1 []string, words2 []string) []string {
	result := []string{}
	// 统计words2中每个字符出现的次数
	letterMap := make(map[string]uint, 26)
	for _, word := range words2 {
		letterMap = mergeWordToCountMap(word, letterMap)
	}

	for _, word := range words1 {
		var flag bool
		tLetterMap := copyCountMap(letterMap)
		for _, c := range word {
			if tLetterMap[string(c)] > 0 {
				tLetterMap[string(c)]--
			}
		}

		for _, v := range tLetterMap {
			if v > 0 {
				// 残留字母，则表示不是
				flag = true
				break
			}
		}

		if !flag {
			result = append(result, word)
		}

	}

	return result
}

func mergeWordToCountMap(word string, countMap map[string]uint) map[string]uint {
	tMap := make(map[string]uint, 26)
	for _, c := range word {
		tMap[string(c)]++
	}

	for k, v := range tMap {
		if val, exist := countMap[k]; exist {
			// 如果已存在，则判断大小，取大的
			if val < v {
				countMap[k] = v
			}
		} else {
			countMap[k] = v
		}
	}

	return countMap
}

func copyCountMap(countMap map[string]uint) map[string]uint {
	result := make(map[string]uint, 26)
	for k, v := range countMap {
		result[k] = v
	}

	return result
}

// @lc code=end
