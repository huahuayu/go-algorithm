package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []int{1, 5, 2, 4, 3, 7, 9, 6}

	sort.Ints(items)

	fmt.Println(interpolationSearch(5, items))
}

func interpolationSearch(key int, a []int) bool {

	low := 0
	high := len(a) - 1

	for low <= high {
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := size * (key - a[low]) / (a[high] - a[low])
			guess = low + offset
		}

		if a[guess] < key {
			low = guess + 1
		} else if a[guess] > key {
			high = guess - 1
		} else {
			return true
		}

	}

	return false
}
