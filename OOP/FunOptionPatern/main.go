package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
	Height int
}

type PersonOptions func(*Person)


func main() {

	person1 := NewPerson(SetName("Ali"), SetAge(26), SetGender("Male"), SetHeight(180))
	fmt.Println(person1)

}

func NewPerson(opts ...PersonOptions) *Person {
	p := &Person{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func SetName(name string) PersonOptions {
	return func(p *Person) {
		p.Name = name
	}
}

func SetAge(age int) PersonOptions {
	return func(p *Person) {
		p.Age = age
	}
}

func SetGender(gender string) PersonOptions {
	return func(p *Person) {
		p.Gender = gender
	}
}

func SetHeight(height int) PersonOptions {
	return func(p *Person) {
		p.Height = height
	}
}
