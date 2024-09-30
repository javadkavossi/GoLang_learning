package main

import (
	"fmt"
)

func main() {

	x := 1

	// -------------------- else if
	if x == 1 {
		fmt.Println("x = 1")
	} else if x == 2 {
		fmt.Println("x = 2")
	} else {
		fmt.Println("x = 3")
	}

	// --------------------

	if x > 5 {
		fmt.Println("x > 5")
	} else {
		fmt.Println("x <= 5")
	}

	// -------------------- AND

	if x > 5 && x < 10 {
		fmt.Println("x > 5 && x < 10")
	}

	// ------------------ OR
	if x > 5 || x < 10 {
		fmt.Println("x > 5 || x < 10")
	}

	// ----------------- fraction

	a := 1.5
	b := 2.5

	if frac := a / b; frac > 0.5 {
		fmt.Println("frac > 0.5")
	}

	switch_statement()

}


func switch_statement() {

	x := 2
	switch x {
	case 1:
		fmt.Println("x = 1")
	case 2:
		fmt.Println("x = 2")
	default:
		fmt.Println("x = 3")
	}
	
}
