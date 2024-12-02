package main

func main() {
	printHello()
	str2 := printHello2()
	println(str2)

	str3 := printHello3("Hello 3")
	println(str3)
}


func printHello() {
	println("Hello")
}


func printHello2() string {
	return "Hello 2"
}

func printHello3(str string) string {
	return str
}



