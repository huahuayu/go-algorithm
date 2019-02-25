package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := generateSlice(20)
	fmt.Println("origin slice is: ", a)

	bubbleSort(a)
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

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
