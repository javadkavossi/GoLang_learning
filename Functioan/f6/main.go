package main

import "time"

func main() {

	name := []string{"javad", "ali", "mohammad", "reza", "sara", "javad", "ali", "mohammad", "reza", "sara"}

	for i, name := range name {
		go func() {
			name = "*" + name
			println(i, name)
		}()
	}
	time.Sleep(time.Second * 1)

	//------------------------------------------

	for i, name := range name {
		go func(index int, item string) {
			name = "*" + name
			println(index, item)
		}(i, name)
	}
	time.Sleep(time.Second * 1)
}
