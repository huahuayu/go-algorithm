package solution

import (
	"math/rand"
	"testing"
	"time"
)

func TestSplitArrary(t *testing.T) {
	s1 := []int{6, 4, -3, 5, -2, -1, 0, 1, -9} // original input
	s2 := generateSlice(20)                    // random input
	var s3 []int                               // border condition1
	s4 := []int{1, 2, 3, 4}                    // border condition2
	s5 := []int{-1, -2, -3, -4}                // border condition3

	splitArray(s1)
	splitArray(s2)
	splitArray(s3)
	splitArray(s4)
	splitArray(s5)

	if !isValidate(s1) {
		t.Error("s1 not pass", s1)
	}

	if !isValidate(s2) {
		t.Error("s2 not pass", s2)
	}

	if !isValidate(s3) {
		t.Error("s3 not pass", s3)
	}

	if !isValidate(s4) {
		t.Error("s4 not pass", s4)
	}

	if !isValidate(s5) {
		t.Error("s5 not pass", s5)
	}
}

// generate random num from -100 ~ 100
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(200) - 100
	}
	return slice
}

// validate slice: positive in left && negative right, true; otherwise false
func isValidate(a []int) bool {
	if a == nil {
		return true
	}

	checkNegative := false

	for i, j := 0, 1; j < len(a); j++ { // i慢指针，j快指针
		if i < 0 && j <= 0 {
			return false
		}

		if i >= 0 && j < 0 {
			checkNegative = true
		}

		if checkNegative == true && a[j] > 0 {
			return false
		}

	}

	return true

}
