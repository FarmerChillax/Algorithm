/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode.cn/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (66.22%)
 * Likes:    973
 * Dislikes: 0
 * Total Accepted:    331.7K
 * Total Submissions: 500.8K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1,3,4]
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1,null,3]
 * 输出: [1,3]
 *
 *
 * 示例 3:
 *
 *
 * 输入: []
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * 二叉树的节点个数的范围是 [0,100]
 * -100
 *
 *
 */
package algorithm

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	secStack := []*TreeNode{}
	ans := []int{}
	for index := 0; index < len(stack); index++ {
		node := stack[index]
		if node.Left != nil {
			secStack = append(secStack, node.Left)
		}

		if node.Right != nil {
			secStack = append(secStack, node.Right)
		}

		if index == len(stack)-1 {
			// 当前层最后一个元素
			ans = append(ans, node.Val)
			stack = secStack
			secStack = []*TreeNode{}
			index = -1
		}
	}
	return ans
}

// @lc code=end
