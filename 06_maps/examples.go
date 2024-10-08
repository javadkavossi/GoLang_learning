package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
}

func main() {

	persons := make(map[string]Person2)

	persons["123"] = Person2{Name: "Ali", Age: 12}
	persons["124"] = Person2{Name: "Ali2", Age: 14}

	fmt.Println(persons)

	persons["123"] = Person2{Name: "update", Age: 21}
	fmt.Println(persons)

	delete(persons, "124")

	test, ok := persons["124"]

	if ok {
		fmt.Println(test.Name)
	} else {
		fmt.Println("not found")
	}

}
