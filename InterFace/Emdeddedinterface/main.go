package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
	Walk()
}

type Human interface {
	Animal
	Speak()
	Think()
}

type Employee struct {
	Name string
	Age  int
}

type Cat struct {
	Name string
}

func main() {

	employee := &Employee{
		Name: "John",
		Age:  30,
	}

	cat := &Cat{
		Name: "Cat",
	}

	var human Human = employee
	var animal Animal = cat

	human.Eat()
	human.Sleep()
	human.Walk()
	human.Think()
	human.Speak()

	animal.Eat()
	animal.Sleep()
	animal.Walk()

}

func (cat *Cat) Eat() {
	fmt.Println("Cat is Eating")
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is Sleeping")
}

func (cat *Cat) Walk() {
	fmt.Println("Cat is Walking")
}

func (employee *Employee) Eat() {
	fmt.Println("Employee is Eating")
}

func (employee *Employee) Sleep() {
	fmt.Println("Employee is Sleeping")
}

func (employee *Employee) Walk() {
	fmt.Println("Employee is Walking")
}

func (employee *Employee) Think() {
	fmt.Println("Employee is Thinking")
}

func (employee Employee) Speak() {
	fmt.Println("Employee is Speaking")
}
