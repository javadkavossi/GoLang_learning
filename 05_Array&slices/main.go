package main

import "fmt"

func main() {
	var myArr0 [5]int
	var myArr1 [2]int = [2]int{1, 2}
	myArr2 := [2]int{1, 2}
	myArr3 := [...]int{1, 2}

	fmt.Println(myArr0, "\n", myArr1, "\n", myArr2, "\n", myArr3)

}
