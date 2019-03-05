package solution

import (
	. "algorithm/leetcode/common/linkedlist"
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	fmt.Println(head)
	fmt.Println(reverseList(head))
}

func TestReverseListRecursive(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseListRecursive(head))
}
