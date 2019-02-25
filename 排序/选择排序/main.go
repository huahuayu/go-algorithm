package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := generateSlice(20)
	fmt.Println("origin slice is: ", a)

	selectionSort(a)
	fmt.Println("sorted slice is:", a)
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}

}
