package main

import (
	"fmt"
	"strings"
)

func main() {

	names := [4]string{"ali", "mohamad", "javad"}

	for _, item := range names {
		item = strings.ToUpper(item)
	}

	fmt.Println(names)

	for i, _ := range names {
		names[i] = strings.ToUpper(names[i])
	}

	fmt.Println(names)

}
