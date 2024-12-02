package main

import "fmt"

func main() {

	roomType := "single"

	price, discount, total := CalculateRoomPrice(roomType, 3, 2)

	fmt.Printf("price: %d\n", price)
	fmt.Printf("discount: %f\n", discount)
	fmt.Printf("discount: %f\n", total)

}

func CalculateRoomPrice(roomType string, nighs int, personCount int) (int, float32, float32) {
	var price int
	var discount float32
	switch roomType {
	case "single":
		price = 100 * nighs
	case "double":
		price = 200 * nighs
	case "suite":
		price = 300 * nighs
	}
	discount = float32(price) * 0.3
	total := float32(price) + discount
	return price, discount, total
}


