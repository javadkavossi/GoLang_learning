package main

import (
	"fmt"
	"os"
// 	"github.com/pkg/errors"
 )

type Trade struct {
	// if property name structs with capital letter , that's public otherwise is private :
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// contractor
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol can't be empty")

	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return trade, nil
}

func (trade *Trade) Value() float64 {
	value := float64(trade.Volume) * trade.Price

	if trade.Buy {
		value = -value
	}
	return value
}

func main() {
	// t := Trade{
	// 	"apple",
	// 	10,
	// 	99.2,
	// 	true,
	// }

	t, err := NewTrade("microsoft", 10.0, 100.00, true)
	if err != nil {
		// return errors.Wrap(err ,"Error can't create trad")
		fmt.Printf("Error can't create trade - %s\n ", err)
		os.Exit(1)

	}

	fmt.Println(t.Value())
}
