/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 *
 * https://leetcode.cn/problems/4sum-ii/description/
 *
 * algorithms
 * Medium (64.25%)
 * Likes:    918
 * Dislikes: 0
 * Total Accepted:    232K
 * Total Submissions: 361.2K
 * Testcase Example:  '[1,2]\n[-2,-1]\n[-1,2]\n[0,2]'
 *
 * 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l)
 * 能满足：
 *
 *
 * 0 <= i, j, k, l < n
 * nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
 * 输出：2
 * 解释：
 * 两个元组如下：
 * 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) +
 * (-1) + 2 = 0
 * 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) +
 * (-1) + 0 = 0
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums1.length
 * n == nums2.length
 * n == nums3.length
 * n == nums4.length
 * 1 <= n <= 200
 * -2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28
 *
 *
 */
package algorithm

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{-2, -1},
				nums3: []int{-1, 2},
				nums4: []int{0, 2},
			},
			want: 2,
		},
		{
			name: "case-2",
			args: args{
				nums1: []int{0},
				nums2: []int{0},
				nums3: []int{0},
				nums4: []int{0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.nums1, tt.args.nums2, tt.args.nums3, tt.args.nums4); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
