/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (65.96%)
 * Likes:    1815
 * Dislikes: 0
 * Total Accepted:    886.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[9,20],[15,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1]
 * 输出：[[1]]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 2000] 内
 * -1000 <= Node.val <= 1000
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
func levelOrder(root *TreeNode) [][]int {
	stack := []*TreeNode{root}
	result := [][]int{}
	if root == nil {
		return result
	}
	for len(stack) != 0 {
		tempStack := []*TreeNode{}
		var vals []int
		for _, node := range stack {
			if node.Left != nil {
				tempStack = append(tempStack, node.Left)
			}
			if node.Right != nil {
				tempStack = append(tempStack, node.Right)
			}
			vals = append(vals, node.Val)
		}
		result = append(result, vals)
		stack = tempStack
	}
	return result
}

// @lc code=end
