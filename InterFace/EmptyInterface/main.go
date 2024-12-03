package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	Print("hello")
	Print(1234)
	Print(true)
	Print(Person{Name: "Alireza", Age: 20})

}

func Print(input interface{}) {
	switch input.(type) {
	case string:
		fmt.Printf("%T: %v\n", input, input)
	case int:
		fmt.Printf("%T: %v\n", input, input)
	case bool:
		fmt.Printf("%T: %v\n", input, input)
	case Person:
		fmt.Printf("%T: %v\n", input, input)
	}

}
