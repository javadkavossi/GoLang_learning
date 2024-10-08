package main

import "fmt"

func main() {

	number := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}

	number = append(number, 0)

	count := copy(number[7:], number[6:])
	//1 ,2,3,4,5,6,7,9,10,0
	//1,2,3,4,5,6,7,7,8,9,10
	number[7] = 8
	fmt.Println(number)
	fmt.Println(count)

	//	remove

	number = append(number[:7], number[8:]...)
	fmt.Println(number)
	number = number[:0]

	fmt.Println("len : ", len(number))
	fmt.Println("cap : ", cap(number))
	number = nil

	fmt.Println("len : ", len(number))
	fmt.Println("cap : ", cap(number))
	fmt.Println(number)

}
