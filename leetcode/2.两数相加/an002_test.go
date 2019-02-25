package problem002

import (
	"fmt"
	"testing"
)

var l1, l2 *ListNode

//func init() {
//	l1 = new(ListNode)
//	l2 = new(ListNode)
//	l1.Val = 2
//	l1.Next = new(ListNode)
//	l1.Next.Val = 4
//	l1.Next.Next = new(ListNode)
//	l1.Next.Next.Val = 3
//
//	l2.Val = 5
//	l2.Next = new(ListNode)
//	l2.Next.Val = 6
//	l2.Next.Next = new(ListNode)
//	l2.Next.Next.Val = 4
//
//}

func init() {
	l1 = &ListNode{0, nil}
	l2 = &ListNode{0, nil}
}
func TestNormalCase(t *testing.T) {
	l3 := addTwoNumbers(l1, l2)
	if l3 != nil {
		fmt.Println(l3.Val)
	}

	if l3.Next != nil {
		fmt.Println(l3.Next.Val)
	}

	if l3.Next != nil && l3.Next.Next != nil {
		fmt.Println(l3.Next.Next.Val)
	}

}

func TestSliceToList(t *testing.T) {
	is := []int{6, 7, 5, 8}
	head := sliceToList(is)
	fmt.Println(head.Val, head.Next.Val, head.Next.Next.Val, head.Next.Next.Next.Val)
}

func TestListToSlice(t *testing.T) {
	a := []int{6, 7, 5, 8}
	head := sliceToList(a)
	b := listToSlice(head)

	if len(a) != len(b) {
		t.Error("len not the same")
	}
	for i, v := range a {
		if v != b[i] {
			t.Error("element not the same")
		}
	}

}

func TestListPrint(t *testing.T) {
	a := []int{6, 7, 5, 8}
	head := sliceToList(a)
	fmt.Println(head)
	fmt.Println(a)
}

func TestAddTwoNumbersOfficial(t *testing.T) {
	l1 := sliceToList([]int{6, 7, 5, 8})
	l2 := sliceToList([]int{8, 4, 1, 9})
	l3 := addTwoNumbersOfficial(l1, l2)

	fmt.Println(l3)
}
