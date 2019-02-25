package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []int{1, 5, 3, 7, 9, 2, 6, 8}

	// sort
	sort.Ints(items)

	// binary search
	fmt.Println(binarySearch(2, items))

}

func binarySearch(key int, a []int) bool {

	low := 0
	high := len(a) - 1

	for low <= high {
		mid := low + (high-low)/2

		if a[mid] < key {
			low = mid + 1
		} else if a[mid] > key {
			high = mid - 1
		} else {
			return true
		}
	}

	return false
}
