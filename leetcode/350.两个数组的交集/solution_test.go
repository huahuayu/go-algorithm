package solution

import (
	"fmt"
	"testing"
)

func TestIntersect(test *testing.T) {
	result := intersect([]int{1, 2, 2, 3}, []int{4, 2, 2, 5, 3, 6, 1, 8})

	fmt.Println(result)
}

func TestIntersectWinner(test *testing.T) {
	result := intersectWinner([]int{1, 2, 2, 3}, []int{4, 2, 2, 5, 3, 6, 1, 8})
	fmt.Println(result)
}
