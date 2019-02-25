package solution

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	is1 := []int{1, 0, 2, 4, 0, 3, 5, 0}
	moveZeroes(is1)
	if !reflect.DeepEqual(is1, []int{1, 2, 4, 3, 5, 0, 0, 0}) {
		t.Error("expected {1,2,4,3,5,0,0,0} but got:", is1)
	}

	is2 := []int{0, 0, 1}
	moveZeroes(is2)
	if !reflect.DeepEqual(is2, []int{1, 0, 0}) {
		t.Error("expected {1,0,0} but got:", is2)
	}

	is3 := []int{0, 1, 0, 3, 12}
	moveZeroes(is3)
	if !reflect.DeepEqual(is3, []int{1, 3, 12, 0, 0}) {
		t.Error("expected {1, 3, 12, 0, 0} but got:", is3)
	}

}
