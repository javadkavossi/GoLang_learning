package main

import "fmt"

type Trade struct {
	// if property name structs with capital letter , that's public otherwise is private :
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{
		"apple",
		10,
		99.2,
		true,
	}
	fmt.Println(t1)
	fmt.Printf("%+v\n",t1)
	fmt.Println(t1.Symbol)
}
