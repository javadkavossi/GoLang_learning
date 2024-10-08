package main

import "fmt"

func main() {

	number := [4]int{1, 2, 3, 4}

	number2 := number
	number3 := &number

	number2[1] = 100
	number3[1] = 200

	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)

	changeValue(&number)
	fmt.Println(number)

	changeValue2(number)
	fmt.Println(number)
}

func changeValue(myArray *[4]int) {
	myArray[0] = 101
	myArray[1] = 103

}

func changeValue2(myArray [4]int) {
	myArray[0] = 150
	myArray[1] = 160
	fmt.Println(len(myArray), myArray)
}
