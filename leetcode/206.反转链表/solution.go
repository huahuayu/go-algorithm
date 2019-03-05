package solution

import (
	. "algorithm/leetcode/common/linkedlist"
)

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

/*
方法一：迭代
var reverseList = function(head) {
    if(head == null) {
        return null
    }
    let p = head
    let newH = head
    while(head.next != null) {
        p = head.next
        head.next = p.next
        p.next = newH
        newH = p
    }
    return newH
};

方法二：递归
let m_phead = null
var reverse = function(head) {
    if(head == null) {
        return null
    }
    if(head.next == null) {
        m_phead = head
        return head
    }
    let tail = reverse(head.next)
    tail.next = head
    head.next = null
    return head
};

var reverseList = function(head) {
    m_phead = null
    reverse(head)
    return m_phead
}
*/
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var prev *ListNode // 当前节点的前一个节点
	curr := head       // 当前节点

	for curr != nil { // 循环条件：当前节点不等于nil，如果当前节点 == nil说明已经遍历到了最后，则当前节点的前一个节点就是新的头节点
		next := curr.Next // 暂存下当前节点的下一个节点
		curr.Next = prev  // 当前节点的next反向指向为prev
		prev = curr       // 为下一次循环初始化prev，新的prev节点为current节点
		curr = next       // 新的current节点为next节点
	}

	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {

	if head == nil || head.Next == nil { // 基础case
		return head
	} else {
		newHead := reverseListRecursive(head.Next) // 从第二个节点开始反转，从第二个节点到最后一个节点全反转好了
		head.Next.Next = head                      // 第二个节点反向指向第一个节点
		head.Next = nil                            // 第一个节点指向nil
		return newHead
	}
}
