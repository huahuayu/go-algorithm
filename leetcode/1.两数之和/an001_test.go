package problem001

import (
	"fmt"
	"testing"
)

var is []int = []int{7, 2, 3, 3, 5, 6, 8, 9, 10}
var isNegative []int = []int{3, -1, 7, 2, 11, 2, 5, 6, 8, 9, 10}

// 普通案例，两个不相同的值
func TestTwoSum(t *testing.T) {
	target := 7
	if sub := twoSum(is, target); sub != nil {
		fmt.Println(sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer1")
		}
	}

	if sub := twoSumHash(is, target); sub != nil {
		fmt.Println(sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer2")
		}
	}

	if sub := twoSumHashImproved(is, target); sub != nil {
		fmt.Println(sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer3")
		}
	}

}

// 测两个相同元素值
func TestTwoSameElements(t *testing.T) {
	target := 6

	if sub := twoSum(is, target); sub != nil {
		fmt.Println(sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer1")
		}
	}

	if sub := twoSumHash(is, target); sub != nil {
		fmt.Println(sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer2")
		}
	}

	if sub := twoSumHashImproved(is, target); sub != nil {
		fmt.Println(sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer3")
		}
	}
}

// 测无匹配
func TestNil(t *testing.T) {
	target := 4

	if sub := twoSum(is, target); sub != nil {
		fmt.Println("1", sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer1")
		}
	}

	if sub := twoSumHash(is, target); sub != nil {
		fmt.Println("2", sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer2")
		}
	}

	if sub := twoSumHashImproved(is, target); sub != nil {
		fmt.Println("3", sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer3")
		}
	}
}

// 测负数
func TestNegative(t *testing.T) {
	target := 6
	is = isNegative

	if sub := twoSum(is, target); sub != nil {
		fmt.Println("1", sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer1")
		}
	}

	if sub := twoSumHash(is, target); sub != nil {
		fmt.Println("2", sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer2")
		}
	}

	if sub := twoSumHashImproved(is, target); sub != nil {
		fmt.Println("3", sub)
		if is[sub[0]]+is[sub[1]] != target {
			t.Error("not the answer3")
		}
	}
}
