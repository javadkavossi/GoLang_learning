package main

import "fmt"

func main() {

	i, j := 1, 2

	var ip *int = &i
	var jp *int = &j

	fmt.Println("Address of i:", &i)
	fmt.Println("Address of ip:", ip)

	fmt.Println("Address of j:", &j)
	fmt.Println("Address of jp:", jp)

	i2 := i
	i2 = i2 + 2

	fmt.Println("Address of i2:", i2)
	fmt.Println("Address of i:", i)

	var ip2 *int = &i
	*ip2 = *ip2 + 2
	fmt.Println("Address of ip2:", ip2)
	fmt.Println("Address of i:", i)

	fmt.Println("======================================================")

	num := 20
	println("main Address ", &num)
	changeByValue(num)
	fmt.Println("changeByValue in main :", num)

	println("==================")

	changeNumberByRef(&num)
	fmt.Println("changeNumberByRef in main :", num)

}

func changeNumberByRef(num *int) {
	*num = *num + 2
	fmt.Println("changeNumberByRef:", num)
	fmt.Println("changeNumberByRef Address :", num)
}

func changeByValue(num int) {
	num = num + 2
	fmt.Println("changeByValue :", num)
	fmt.Println("changeByValue  Address:", &num)

}
