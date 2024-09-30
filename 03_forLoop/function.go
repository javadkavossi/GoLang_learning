package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------- break -----------")

	for i := 0; i < 5; i++ {

		if i == 3 {
			break
		}
		fmt.Println("print i and break at i =", i)
	}

	fmt.Println("----------- continue -----------")

	for i := 0; i < 5; i++ {

		if i <= 3 {
			continue
		}
		fmt.Println("print i and continue at i =", i)
	}

}
