package main

import "fmt"

func safeValues(values []int, index int)int {
	defer func() {
		if err := recover(); err != nil {

			fmt.Printf("Error %s\n", err)
		}

	}()

	return values[index]
}

func main() {

	// file, err := os.Open("no-file")
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()
	// fmt.Println("file opening ... ")
	v := safeValues([]int{1, 2, 3}, 5)
	fmt.Println(v)
}
