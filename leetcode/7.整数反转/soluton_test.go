package solution

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	x1 := 123456
	x2 := -123456
	x3 := 2147483641
	x4 := 2147483647
	x5 := 2147483648
	x6 := -2147483647
	x7 := -2147483648
	x8 := -2147483649
	fmt.Println(reverse(x1))
	fmt.Println(reverse(x2))
	fmt.Println(reverse(x3))
	fmt.Println(reverse(x4))
	fmt.Println(reverse(x5))
	fmt.Println(reverse(x6))
	fmt.Println(reverse(x7))
	fmt.Println(reverse(x8))
}
