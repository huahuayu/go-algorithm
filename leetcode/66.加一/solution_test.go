package solution

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	result := plusOne([]int{1, 2, 3})
	if !reflect.DeepEqual(result, []int{1, 2, 4}) {
		t.Error("expected [1,2,4] but got:", result)
	}

	result = plusOne([]int{9, 9, 9})
	if !reflect.DeepEqual(result, []int{1, 0, 0, 0}) {
		t.Error("expected [1,0,0,0] but got:", result)
	}
}
