package main

import (
	"fmt"
	"time"
)

//  We have 2 kinds of channels : (Buffered , )


func main() {

	channel := make(chan int)

	// Goroutine
	go func() {
		channel <- 25 // Push To Channel
	}()

	val := <-channel
	fmt.Printf("go %d\n", val)
	fmt.Println("----------------")

	//===================================  Send Multiple

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			channel <- i
			time.Sleep(1 * time.Second) // 1 second
		}
	}()

	//===================================  Receive Multiple

	for i := 0; i < 3; i++ {
		val := <-channel
		fmt.Printf("received %d\n", val)
	}

	//===================================  Close to Signal we are Done ///

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("sending %d\n", i)
			channel <- i
			time.Sleep(1 * time.Second) // 1 second
		}
		close(channel) // Close Channel
		fmt.Println("Close to Signal channel")
	}()

	for i := range channel {
		fmt.Printf("received  %d\n", i)
	}

}
