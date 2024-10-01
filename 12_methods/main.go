package main

import "fmt"

type Trade struct {
	// if property name structs with capital letter , that's public otherwise is private :
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price

	if trade.Buy {
		value = -value
	}
	return value
}

func main() {
	t := Trade{
		"apple",
		10,
		99.2,
		true,
	}
	fmt.Println(t.Value())

}
