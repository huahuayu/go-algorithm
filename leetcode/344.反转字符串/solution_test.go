package solution

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "abcdefg"
	bs := []byte(s)
	fmt.Println("original bs:", bs)
	reverseString(bs)
	fmt.Println("reversed bs:", string(bs))

}
