package main

import "fmt"

func main() {

	apiResponse := struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}{
		Success: true,
		Data: struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			Name: "Alireza",
			Age:  20,
		},
	}

	fmt.Printf("apiResponse: %v\n", apiResponse)

}
