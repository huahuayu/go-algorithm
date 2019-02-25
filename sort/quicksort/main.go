package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := generateSlice(20)
	fmt.Println("origin slice is: ", a)

	quickSort(a)
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

func quickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	mid := a[0]
	head, tail := 0, len(a)-1
	for i := 1; i <= tail; {
		if a[i] > mid {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[i], a[head] = a[head], a[i]
			head++
			i++
		}
		fmt.Println(head, i)
	}
	quickSort(a[:head])
	quickSort(a[head+1:])
}
