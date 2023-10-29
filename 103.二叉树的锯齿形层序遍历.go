/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 *
 * https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (57.77%)
 * Likes:    827
 * Dislikes: 0
 * Total Accepted:    327.6K
 * Total Submissions: 566.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：[[3],[20,9],[15,7]]
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

var (
	zigzagRes [][]int
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	zigzagRes = make([][]int, 0)
	if root == nil {
		return zigzagRes
	}
	leftToRight([]*TreeNode{root})
	return zigzagRes
}

func leftToRight(nodes []*TreeNode) {
	if len(nodes) == 0 {
		return
	}
	nextLevel := []*TreeNode{}
	vals := make([]int, 0, len(nodes))
	for i := 0; i < len(nodes); i++ {
		vals = append(vals, nodes[i].Val)
		if nodes[i].Left != nil {
			nextLevel = append(nextLevel, nodes[i].Left)
		}
		if nodes[i].Right != nil {
			nextLevel = append(nextLevel, nodes[i].Right)
		}
	}
	zigzagRes = append(zigzagRes, vals)
	rightToLeft(nextLevel)
}

func rightToLeft(nodes []*TreeNode) {
	if len(nodes) == 0 {
		return
	}
	nextLevel := []*TreeNode{}
	vals := make([]int, 0, len(nodes))
	for i := len(nodes) - 1; i >= 0; i-- {
		vals = append(vals, nodes[i].Val)
		if nodes[i].Right != nil {
			nextLevel = append([]*TreeNode{nodes[i].Right}, nextLevel...)
		}
		if nodes[i].Left != nil {
			nextLevel = append([]*TreeNode{nodes[i].Left}, nextLevel...)
		}

	}
	zigzagRes = append(zigzagRes, vals)
	leftToRight(nextLevel)
}

// @lc code=end
