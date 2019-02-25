package solution

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	result := singleNumber([]int{1, 2, 2, 1, 3, 4, 4})
	fmt.Println(result)
	fmt.Println("1^1", 1^1)
	fmt.Println("1^2", 1^2)
	fmt.Println("3^6^6", 3^6^6)

	a, b := 5, 10

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a, b)

}
