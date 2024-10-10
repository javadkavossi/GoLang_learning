package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type PersonsOptions struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	ali := Person{FirstName: "John", LastName: "Doe", Age: 42}

	fmt.Println(ali)

	ali2 := new(Person)
	ali2.FirstName = "Jane"
	ali2.LastName = "Doe"
	ali2.Age = 12

	fmt.Println(ali2)

	ali3 := &Person{FirstName: "John", LastName: "Doe", Age: 42}
	fmt.Println(ali3)

	ali4 := NewPerson("ali", "mohammedan", 12)
	fmt.Println(ali4)

	ali5 := NewPersonOptions(PersonsOptions{"ahmad", "abort 2", 23})

	fmt.Println(ali5)

}

func NewPerson(firsName, lastName string, age int) *Person {

	return &Person{FirstName: firsName, LastName: lastName, Age: age}

}

func NewPersonOptions(personsOptions PersonsOptions) *Person {
	return &Person{FirstName: personsOptions.LastName, LastName: personsOptions.FirstName, Age: personsOptions.Age}
}
