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

	fmt.Println("----------- nested loop -----------")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
	}

	fmt.Println("----------- nested loop break -----------")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("----------- nested loop continue -----------")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("----------- infinite loop -----------")

	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}

	fmt.Println("----------- will -----------")
	b := 0

	for {
		fmt.Println(b)
		b++
		if b == 5 {
			break
		}
	}

}
