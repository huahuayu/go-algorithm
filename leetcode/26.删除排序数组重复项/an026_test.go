package problem026

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	is := []int{1, 2, 2, 2, 3, 5, 7, 8, 8, 9}
	fmt.Println(removeDuplicates(is))
	fmt.Println(is)

}
