package solution

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	s1 := "leetcode"
	s2 := "loveleetcode"
	//r1 := firstUniqChar(s1)
	//r2 := firstUniqChar(s2)
	r3 := firstUniqCharWinner(s1)
	r4 := firstUniqCharWinner(s2)
	fmt.Println(r3, r4)
}
