package main

import "fmt"

func main() {

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	number2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	number = append(number, 11, 12, 13, 14, 15)

	fmt.Println(number)

	number = append(number, number2...)

	fmt.Println(number)

}
