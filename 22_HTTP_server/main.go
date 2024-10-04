package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fprint, err := fmt.Fprintf(w, "Hello World")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fprint)

}

func main() {

	http.HandleFunc("/", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
