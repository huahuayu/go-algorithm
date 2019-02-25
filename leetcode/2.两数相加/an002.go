package problem002

import (
	"fmt"
	"math"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力法
/*
1560 / 1563 个通过测试用例
输入：
[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
[5,6,4]
输出：
[-2]
预期：
[6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// ax储存链表的data部分为slice
	var a1, a2 []int

	// nx将slice拼接为int数字
	var n1, n2, n3 int

	// 将l1映射为a1
	for {
		if l1 == nil {
			a1 = []int{0}
			break
		}
		if l1.Next == nil {
			a1 = append(a1, l1.Val)
			break
		}
		a1 = append(a1, l1.Val)
		l1 = l1.Next
	}

	for {
		if l2 == nil {
			a2 = []int{0}
			break
		}
		if l2.Next == nil {
			a2 = append(a2, l2.Val)
			break
		}
		a2 = append(a2, l2.Val)
		l2 = l2.Next
	}

	if a1 == nil {
		n1 = 0
	} else {
		for i, n := range a1 {
			n1 = n1 + n*int(math.Pow(float64(10), float64(i)))
		}
	}

	if a2 == nil {
		n2 = 0
	} else {
		for i, n := range a2 {
			n2 = n2 + n*int(math.Pow(float64(10), float64(i)))
		}
	}

	n3 = n1 + n2
	var n3Slice []int

	for {
		if n3/10 > 0 {
			n3Slice = append(n3Slice, n3%10)
			n3 = n3 / 10
		} else {
			n3Slice = append(n3Slice, n3%10)
			break
		}
	}

	//fmt.Println("n3slice", n3Slice)
	var l3 *ListNode

	for i := 0; i < len(n3Slice); i++ {
		node := ListNode{n3Slice[i], nil}
		if l3 == nil {
			l3 = &node
		} else {
			last := l3
			for {
				if last.Next == nil {
					break
				}
				last = last.Next
			}

			last.Next = &node
		}
	}

	return l3
}

func addTwoNumbersOfficial(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	p := l1
	q := l2
	curr := dummyHead

	var x, y, sum, carry int

	for p != nil || q != nil {

		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum = carry + x + y
		carry = sum / 10 // carry代表进位数字，如果当前位相加有进位则carry会等于1，遗留到下一位使用

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

func sliceToList(is []int) (head *ListNode) {
	if is == nil {
		return nil
	}

	head = new(ListNode)
	curr := head
	for i := 0; i < len(is); i++ {
		curr.Val = is[i]
		if i != len(is)-1 {
			curr.Next = new(ListNode)
			curr = curr.Next
		}
	}
	return head
}

func listToSlice(head *ListNode) (is []int) {
	if head == nil {
		return nil
	}

	curr := head
	for curr != nil {
		is = append(is, curr.Val)
		curr = curr.Next
	}

	return is
}

func (head *ListNode) String() string {
	is := listToSlice(head)
	return fmt.Sprintf("%v", is)
}
