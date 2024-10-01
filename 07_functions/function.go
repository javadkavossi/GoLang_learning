package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divMod(a int, b int) (float64, int) {
	return float64(a) / float64(b), a % b
}

func main() {

	val := add(10, 2)
	fmt.Println(val)

	div, mod := divMod(70, 8)
	fmt.Printf("div = %.2f , mod = %d\n", div, mod)

}
