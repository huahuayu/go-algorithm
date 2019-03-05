package solution

import (
	. "algorithm/leetcode/common/linkedlist"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	l1 := NewListNode([]int{})
	l2 := NewListNode([]int{1})
	l3 := NewListNode([]int{1, 2, 3, 4, 3, 2, 1})
	l4 := NewListNode([]int{1, 2, 3})

	if isPalindrome(l1) == false {
		t.Errorf("expect l1 true, but got %v", isPalindrome(l1))
	}

	if isPalindrome(l2) == false {
		t.Errorf("expect l2 true, but got %v", isPalindrome(l2))
	}

	if isPalindrome(l3) == false {
		t.Errorf("expect l3 true, but got %v", isPalindrome(l3))
	}

	if isPalindrome(l4) == true {
		t.Errorf("expect l4 false, but got %v", isPalindrome(l4))
	}
}
