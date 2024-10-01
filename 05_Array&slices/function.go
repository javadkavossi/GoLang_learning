package main

import (
	"fmt"
)

func main() {
	foo := []string{"ali", "amir", "test"}
	fmt.Printf("foo = %v (type = %T)\n", foo, foo)
	// length
	fmt.Println(len(foo))
	// 0 index
	fmt.Println(foo[1])
	//slices
	fmt.Println(foo[1:])

	// for loop
	for i := 0; i < len(foo); i++ {
		fmt.Println(foo[i])
	}

	// Single value range (to show indexes)

	for index := range foo {
		fmt.Println(index)
	}
	// Double value range (to show index and value together)

	for index, value := range foo {
		fmt.Printf("%d - %s\n", index, value)
	}
	// Double value range ignore index by using _
	for _, value := range foo {
		fmt.Println(value)
	}

	// Append

	foo = append(foo, "test 3")
	println(foo)

}
