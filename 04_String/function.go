package main

import (
	"fmt"
)

func main() {
	book := "GoLang_learning"
	fmt.Println("book = ", book)
	fmt.Println("book[0] = ", book[0])
	fmt.Println("book[5] = ", book[5])
	// ------------------------- Slicing
	fmt.Println("book[0:5] = ", book[3:5])
	// ------------------------- Slicing no end
	fmt.Println("book[5:] = ", book[5:])

	// ------------------------- len
	fmt.Print("len(book) = ", len(book))

	// ------------------------- for loop
	for i := 0; i < len(book); i++ {
		fmt.Println("book[i] = ", book[i])
	}

	// ------------------------- for range
	for index, value := range book {
		fmt.Printf("book[%v] = %v , type of book[%v] = %T\n", index, value, index, value)
	}

	// ------------------------- for range
	fmt.Printf("book[0] =%v , type of book[0] = %T\n", book[0], book[0])

	// ------------------------- Use + to concatenate
	book2 := "Go"
	fmt.Println("book2 = ", book2)
	fmt.Println("book + book2 = ", book+book2)

	// ------------------------- Use += to concatenate
	book += book2
	fmt.Println("book += book2 = ", book)

	// ------------------------- Use to concatenate

	fmt.Println("hey " + book[1:])

	// ------------------------- multiline
	fmt.Println(`
		Hello, World!
		Hello, World!
		Hello, World!
		Hello, World!
		Hello, World!
		Hello, World!
		...
		
		Hello, World!
		Hello, World!
		Hello, World!
		Hello, World!
	`)

}
