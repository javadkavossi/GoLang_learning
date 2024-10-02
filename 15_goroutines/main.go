package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR :%s\n", err)
		return

	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("%s --> %s\n", url, ctype)
}

var wg sync.WaitGroup // Step 1 

func main() {

	urls := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com"}

	for _, url := range urls {
		wg.Add(1) // Step 3 
		go func(url string) {
			returnType(url)
			wg.Done() // End Step 4  
		}(url)
	}
	wg.Wait() // Step 2

}
