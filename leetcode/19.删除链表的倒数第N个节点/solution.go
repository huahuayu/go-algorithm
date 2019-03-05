package solution

import (
	. "algorithm/leetcode/common/linkedlist"
	"log"
)

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

/*
java方法一：遍历两次
public ListNode removeNthFromEnd(ListNode head, int n) {
    ListNode dummy = new ListNode(0);
    dummy.next = head;
    int length  = 0;
    ListNode first = head;
    while (first != null) {
        length++;
        first = first.next;
    }
    length -= n;
    first = dummy;
    while (length > 0) {
        length--;
        first = first.next;
    }
    first.next = first.next.next;
    return dummy.next;
}
*/

/*
 java方法二：遍历一次
 public ListNode removeNthFromEnd(ListNode head, int n) {
    ListNode dummy = new ListNode(0);
    dummy.next = head;
    ListNode first = dummy;
    ListNode second = dummy;
    // Advances first pointer so that the gap between first and second is n nodes apart
    for (int i = 1; i <= n + 1; i++) {
        first = first.next;
    }
    // Move first to the end, maintaining the gap
    while (first != null) {
        first = first.next;
        second = second.next;
    }
    second.next = second.next.next;
    return dummy.next;
}
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	len := 0
	curr := head
	for curr != nil {
		len++
		curr = curr.Next
	}

	len -= n
	curr = dummy
	for len > 0 {
		len--
		curr = curr.Next
	}

	curr.Next = curr.Next.Next
	return dummy.Next

}

func removeNthFromEndWinner(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first, second := dummy, dummy

	// first先移动n+1步
	for i := 0; i <= n; i++ {
		if first == nil {
			log.Fatalf("invald n=%d greater than len(head)", n)
		}
		first = first.Next
	}

	// first - second = n 步，保持这个距离
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next

}
