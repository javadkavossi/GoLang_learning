package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatalf("http.Get err: %v", err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

	fmt.Println("===============")

	// POST request

	
	job := &Job{
		User:   "Amir",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer // in-memory io reader
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("ERROR: can't encode job - %s", err)
	}

	res, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("ERROR: can't call httpbin.org")
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
