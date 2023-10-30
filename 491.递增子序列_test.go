/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode.cn/problems/non-decreasing-subsequences/description/
 *
 * algorithms
 * Medium (51.88%)
 * Likes:    737
 * Dislikes: 0
 * Total Accepted:    155.5K
 * Total Submissions: 300K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
 *
 * 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,6,7,7]
 * 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,4,3,2,1]
 * 输出：[[4,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 15
 * -100 <= nums[i] <= 100
 *
 *
 */
package algorithm

import (
	"reflect"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				nums: []int{4, 6, 7, 7},
			},
			want: [][]int{
				{4, 6},
				{4, 6, 7},
				{4, 6, 7, 7},
				{4, 7},
				{4, 7, 7},
				{6, 7},
				{6, 7, 7},
				{7, 7},
			},
		},
		{
			name: "case-2",
			args: args{
				nums: []int{4, 4, 3, 2, 1},
			},
			want: [][]int{
				{4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
