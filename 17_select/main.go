package main

import (
	"fmt"
	"time"
)

//  We have 2 kinds of channels : (Buffered , )

func main() {

	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch2 <- 12
	}()

	select {
	case val := <-ch1:
		fmt.Printf("got ch1: %v\n", val)
	case val := <-ch2:
		fmt.Printf("got ch2: %v\n", val)
	}

	fmt.Println("================")

	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Microsecond)
		out <- 1.2
	}()

	select {
	case val := <-out:
		fmt.Printf("got %v\n", val)
	case <-time.After(20 * time.Microsecond):
		fmt.Println("timeout")

	}
}
