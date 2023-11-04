/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode.cn/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (55.81%)
 * Likes:    1678
 * Dislikes: 0
 * Total Accepted:    446.3K
 * Total Submissions: 799.7K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目为 n
 * 1
 * -500
 * 1
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
 *
 */
package algorithm

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 设置链表哨兵
	sentry := &ListNode{Next: head}
	// 前指针
	pre := sentry
	// 移动前驱指针指向 left 节点的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// 移动后继指针指向 right 节点
	rightPointer := pre
	for i := 0; i < right-left+1; i++ {
		rightPointer = rightPointer.Next
	}

	// 切割子链表
	leftNode := pre.Next
	curr := rightPointer.Next
	// 切断链接
	pre.Next = nil
	rightPointer.Next = nil
	// 开始翻转链表
	var tmpPre *ListNode
	cur := leftNode
	for cur != nil {
		next := cur.Next
		cur.Next = tmpPre
		tmpPre = cur
		cur = next
	}

	// 拼接链表
	pre.Next = rightPointer
	leftNode.Next = curr
	return sentry.Next
}

// @lc code=end
