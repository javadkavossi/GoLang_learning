package main

import (
	"fmt"
)

func main() {
	bar := map[string]float64{
		"amazon": 123.2,
		"Google": 523.2,
		"paypal": 12.1,
	}

	// Get a Value
	fmt.Println(bar["amazon"])
	//Get zero value if not found
	fmt.Println(bar["apple"]) // 0

	//use two value to see if found

	value, ok := bar["apple"]

	if !ok {
		fmt.Println("not found ")

	} else {
		fmt.Println(value)
	}

	// set
	bar["apple"] = 0.63
	fmt.Println(bar)

	// Delete
	delete(bar, "amazon")
	println(bar)

	//Single value "for" is not keys

	for key := range bar {
		fmt.Println(key)
	}

	//Double value "for" is key  value

	for key, value := range bar {
		fmt.Printf("%s - %.2f\n", key, value)
	}

}
