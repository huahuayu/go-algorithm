package main

import "fmt"

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	if index := linearSearch(items, 45); index == nil {
		fmt.Println("not exist")
	} else {
		fmt.Printf("exist, index is %d\n", *index)
	}

}

func linearSearch(a []int, key int) (index *int) {
	for i, item := range a {
		if item == key {
			index = &i
			return
		}
	}
	return nil
}
