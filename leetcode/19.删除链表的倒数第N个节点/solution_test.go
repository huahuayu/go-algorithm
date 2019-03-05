package solution

import (
	. "algorithm/leetcode/common/linkedlist"
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(head)
	fmt.Println(removeNthFromEnd(head, 5))
}

func TestRemoveNthFromEndWinner(t *testing.T) {

	head := NewListNode([]int{1})
	fmt.Println(head)
	fmt.Println(removeNthFromEndWinner(head, 1))

	head = NewListNode([]int{1})
	fmt.Println(head)
	fmt.Println(removeNthFromEndWinner(head, 2))

}
