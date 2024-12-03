package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name   string `json:"name" xml:"Name_User"`
	Age    int    `json:"age,omitempty" xml:"Age_User"`
	Gender string `json:"gender" xml:"Gender_User"`
	Height int    `json:"height" xml:"Height_User"`
}

func main() {

	person := Person{Name: "Alireza", Gender: "Male", Height: 180}

	personJson, _ := json.Marshal(person)

	fmt.Println(string(personJson))

	personXml, _ := xml.MarshalIndent(person, "  ", "  ")

	fmt.Println(string(personXml))

}
