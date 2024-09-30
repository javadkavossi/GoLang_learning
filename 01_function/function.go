package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// var x float64
	// var y float64

	// x := 1.0
	// y := 2.0

	x, y := 1.0, 2.0

	fmt.Printf("x = %v , type of x = %T\n", x, x)
	fmt.Printf("y = %v , type of y = %T\n", y, y)

	// var mean float64
	mean := (x + y) / 2
	fmt.Printf("mean = %v , type of mean = %T\n", mean, mean)

}
