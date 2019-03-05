package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	m1 := map[string]interface{}{"A": 1, "B.A": 2, "B.B": "some string", "CC.D.E": "another string", "CC.D.F": 5} // original case
	fmt.Println(unFlatten(m1))

	m2 := map[string]interface{}{} // border case
	fmt.Println(unFlatten(m2))

	m3 := map[string]interface{}{"A.B.C.D.E.F": 2}
	fmt.Println(unFlatten(m3))
}
