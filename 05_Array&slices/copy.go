package main

import "fmt"

func main() {

	number := []int{1, 2, 3, 4, 5}
	number2 := []int{6, 7, 8, 9, 10}

	copy(number, number2)

	fmt.Println(number)
	fmt.Println(number2)
}
