package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
	Height int
}

type PersonBuilder struct {
	Person
}

func main() {

	person := Person{Name: "Alireza", Age: 26, Gender: "Male", Height: 180}

	personBuilder := &PersonBuilder{}

	person2 := personBuilder.SetName("Javad").SetAge(23).SetGender("Male").SetHeight(180).Build()

	fmt.Printf("\n person1: %+v,\n person2: %+v", person, person2)

}

func (builder *PersonBuilder) SetName(name string) *PersonBuilder {
	builder.Name = name
	return builder
}

func (builder *PersonBuilder) SetAge(age int) *PersonBuilder {
	builder.Age = age
	return builder
}

func (builder *PersonBuilder) SetGender(gender string) *PersonBuilder {
	builder.Gender = gender
	return builder
}

func (builder *PersonBuilder) SetHeight(height int) *PersonBuilder {
	builder.Height = height
	return builder
}

func (builder *PersonBuilder) Build() Person {
	return builder.Person
}
