package main

type IPerson interface {
	print()
}

type Person struct {
	Name string
	Age  int
}

func main() {

	var p IPerson

	p = Person{}
	p.print()

}

func (p Person) print() {
	println(p.Name, p.Age)
}
