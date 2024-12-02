package main

import "fmt"

func main() {
	sum, average := Calculate(1, 2, 3, 4, 5)
	println(sum)
	println(average)
	PrintLog(1, "Hello", "World", 1, 2, 3, "Hello", "World")
}

func Calculate(numbers ...int) (sum int, average float32) {
	for _, number := range numbers {
		sum += number
	}
	average = float32(sum) / float32(len(numbers))
	return
}

func PrintLog(message ...interface{}) {
	for i, message := range message {
		fmt.Printf("log %d: %v\n", i, message)
	}
}

