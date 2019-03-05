package solution

import (
	. "algorithm/leetcode/common/linkedlist"
)

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	midNode := findMiddleNode(head)
	rightHalf := reverseLinkList(midNode.Next) // 中间节点之后的部分反转

	p, q := head, rightHalf

	for p != nil && q != nil && p.Val == q.Val { // 比较前半部分和后半部分
		p = p.Next
		q = q.Next
	}

	return q == nil
}

func findMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil { // 快指针一次跳两格，慢指针一次跳一格，快指针停下来时，慢指针处在中间节点上
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseLinkList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

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
