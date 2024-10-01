package main

import "fmt"

func doubleAt(values []int, i int) { //slices , map pass by reference
	values[i] *= 2
}

func double(n int) { // pass by value (int)
	n *= 2
}

func doublePtr(n *int) { //pointer
	*n *= 2
}

func main() {
	values := []int{1, 2, 3 , 4}

	doubleAt(values, 2)

	fmt.Println(values)
	// ===================
	val := 10
	double(val)
	fmt.Println(val)
	// ==================
	doublePtr(&val)

	fmt.Println(val)

}
