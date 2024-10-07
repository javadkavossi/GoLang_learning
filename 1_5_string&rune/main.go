package main

import (
	"fmt"
	"unsafe"
)

func main() {

	myStr := "Hello World 😃 😁"

	fmt.Printf("myStr: %s %T , len %d , size %d \n", myStr, myStr, len(myStr), unsafe.Sizeof(myStr))
	for i := 0; i < len(myStr); i++ {
		fmt.Printf("myStr[%d]: %c %T \n", i, myStr[i], myStr[i])
	}

	fmt.Println("--------------------------------------")

	myRune := []rune(myStr)
	fmt.Printf("myRune: %v %T , len :%d ,size :%d \n", myRune, myRune, len(myRune), unsafe.Sizeof(myRune))
	for i := 0; i < len(myRune); i++ {
		fmt.Printf("myStr[%d]: %c %T\n", i, myRune[i], myRune[i])
	}

}
