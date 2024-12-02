package main

import (
	"fmt"
	"sort"
)

func main() {

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("number: %v\n", number)

	sortingFunction := func(i, j int) bool {
		return number[i] > number[j]
	}
	sort.Slice(number, sortingFunction)
	fmt.Printf("number: %v\n", number)

}
