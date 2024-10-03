package main

import (
	fmt "fmt"
	toml "github.com/pelletier/go-toml"
	log "log"
	os "os"
)

//  We have 2 kinds of channels : (Buffered , )

type config struct {
	Login struct {
		Username string
		Password string
	}
}

func main() {
	file, err := os.Open("./config.toml")
	if err != nil {
		log.Fatalf("open config.toml err:%v", err)
	}
	defer file.Close()
	ctg := &config{}
	dec := toml.NewDecoder(file)

	if err := dec.Decode(ctg); err != nil {
		log.Fatalf("Error: can't decode config file :%v", err)
	}

	fmt.Printf("%+v\n", ctg)
}
