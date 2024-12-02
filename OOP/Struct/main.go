package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	Person1 := Person{
		Name: "Javad",
		Age:  26,
	}

	Person2 := Person{
		Name: "Ali",
		Age:  24,
	}

	fmt.Println(Person1, Person2)
// -----------------------------------------------
	presint4 := new(Person)
	presint4.Age = 25
	presint4.Name = "Reza1"
	fmt.Println(presint4)
// -----------------------------------------------
	Person5 := NewPerson("Reza2", 14)
	fmt.Println(Person5)
// ---------------------------------------------
	NewPerson1 := func(Name string, Age int) *Person {
		return &Person{
			Name: Name,
			Age:  Age,
		}
	}
	Person6 := NewPerson1("Reza6", 26)
	fmt.Println(Person6)
}
func NewPerson(Name string, Age int) *Person {
	if Age < 18 {
		return nil
	}
	return &Person{
		Name: Name,
		Age:  Age,
	}
}
