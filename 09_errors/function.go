package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	}

	return math.Sqrt(n), nil
}

func main() {

	val, err := sqrt(-10.0)
	if err != nil {
		fmt.Println("Error:", err)

	} else {
		fmt.Println("square root : ", val)
	}

	val2, err := sqrt(10.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("square root : ", val2)
	}
}
