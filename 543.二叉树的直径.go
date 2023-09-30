/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 *
 * https://leetcode.cn/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (58.98%)
 * Likes:    1411
 * Dislikes: 0
 * Total Accepted:    348.9K
 * Total Submissions: 591.2K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你一棵二叉树的根节点，返回该树的 直径 。
 *
 * 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
 *
 * 两节点之间路径的 长度 由它们之间边数表示。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,4,5]
 * 输出：3
 * 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [1, 10^4] 内
 * -100 <= Node.val <= 100
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

var result int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLen := getLen(root.Left)
	rightLen := getLen(root.Right)
	result = max(result, leftLen+rightLen)
	return result
}

func getLen(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLen := getLen(root.Left)
	rightLen := getLen(root.Right)
	result = max(result, leftLen+rightLen)
	return max(leftLen, rightLen) + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end
