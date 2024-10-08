/*
	07-04

Send to func and problems
Slice in loop and problems
Copy func
Merge two slice with append
Prepend
Remove and put
Remove all
Remove slice
*/
package main

import "fmt"

func main() {

	myArray := []int{1, 2, 3, 4, 5}

	changeNumber(myArray)
	fmt.Println(myArray)

	addItem(&myArray)
	fmt.Println(myArray)
}

func changeNumber(numbers []int) {

	for i, _ := range numbers {

		numbers[i] = numbers[i] + 2
	}

}

func addItem(numbers *[]int) {
	*numbers = append(*numbers, 6)
}
