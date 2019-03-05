package linkedlist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vals []int) (head *ListNode) {
	if vals == nil {
		return nil
	}

	head = new(ListNode)
	curr := head
	for i, v := range vals {
		curr.Val = v
		if i != len(vals)-1 {
			curr.Next = new(ListNode)
		} else {
			curr.Next = nil
		}
		curr = curr.Next
	}

	return head
}

func (head *ListNode) String() string {
	if head == nil {
		return "nil"
	}

	ss := make([]string, 0)
	for head != nil {
		ss = append(ss, strconv.Itoa(head.Val))
		head = head.Next
	}

	return strings.Join(ss, "->") + "->nil"
}
