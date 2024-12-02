package main

type Employee struct {
	Id            int
	Name          string
	Age           int
	Type          string
	SalaryOfMonth int
}

func main() {

	var employee1 Employee = Employee{Id: 1, Name: "John", Age: 30, Type: "Manager", SalaryOfMonth: 50000}

	println(CalculateSalary(employee1))
	println(employee1.CalculateSalary())
}

// function
func CalculateSalary(employee Employee) int {
	return employee.SalaryOfMonth * 12
}

// method
func (e *Employee) CalculateSalary() int {
	return e.SalaryOfMonth * 12
}
