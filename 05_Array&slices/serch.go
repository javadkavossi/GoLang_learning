package main

import "fmt"

func main() {

	names := [5]string{"ali", "ahamad", "akbar", "sadeg", "javad"}

	Key := "javad"

	for i, name := range names {
		if name == Key {
			fmt.Println("index : ", i)
			break
		}
	}

	//---------------------------------------------------

	for l := 0; l < len(names); l++ {
		if names[l] == Key {
			fmt.Println("nsme found index: ", l)
			break
		}
	}
}
