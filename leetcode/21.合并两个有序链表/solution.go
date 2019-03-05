package solution

import (
	. "algorithm/leetcode/common/linkedlist"
)

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 基础case
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 迭代case
	var res *ListNode
	if l1.Val >= l2.Val { // 取小值节点做head，小值节点的下一个节点做新的头部和l1合并好
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}
