/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode.cn/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (53.62%)
 * Likes:    1818
 * Dislikes: 0
 * Total Accepted:    657K
 * Total Submissions: 1.2M
 * Testcase Example:  '[1,2,2,1]'
 *
 * 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,2,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围[1, 10^5] 内
 * 0 <= Node.val <= 9
 *
 *
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
func isPalindrome(head *ListNode) bool {
	node := head
	var length int
	for ; node != nil; node = node.Next {
		length++
	}
	if length <= 1 {
		return true
	}

	mid := length / 2
	node = head
	for i := 1; i < mid; i++ {
		node = node.Next
	}

	rightListHead := node.Next
	if length%2 == 1 {
		rightListHead = rightListHead.Next
	}

	node.Next = nil
	rightListHead = ReverseList(rightListHead)
	leftListHead := head

	for leftListHead != nil && rightListHead != nil {
		if leftListHead.Val != rightListHead.Val {
			return false
		}
		leftListHead = leftListHead.Next
		rightListHead = rightListHead.Next
	}
	if leftListHead != nil || rightListHead != nil {
		return false
	}
	return true
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// @lc code=end
