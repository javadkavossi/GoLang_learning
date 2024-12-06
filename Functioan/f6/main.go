package main

import "time"

func main() {

	name := []string{"I", "am", "javad", "kavossi", "5", "6", "7", "8", "9", "10"}

	for i, name := range name {
		go func() {
			name = "*" + name
			println(i, name)
		}()
	}
	time.Sleep(time.Second * 1)

	//------------------------------------------

	// for i, name := range name {
	//  func(index int, item string) {
	// 		name = "*" + name
	// 		println(index, item)
	// 	}(i, name)
	// }
	// time.Sleep(time.Second * 1)
}
