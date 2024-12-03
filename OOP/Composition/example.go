package main

import "fmt"

const (
	BaseSalary     = 100
	ExtraHoursRate = 10
	HoursSalary    = 20
	TaxRate        = 0.09
	ShiftSalary    = 30
)

func main() {

	fullTimeEmployee := []FullTimeEmployee{
		{Employee: Employee{Id: 1, Name: "Ali", Age: 30}, ExtraHours: 10.0},
		{Employee: Employee{Id: 2, Name: "John2", Age: 28}, ExtraHours: 5.0},
	}

	partTimeEmployee := []PartTimeEmployee{
		{Employee: Employee{Id: 1, Name: "Akbar", Age: 30}, PartTimeHours: 5.0},
		{Employee: Employee{Id: 2, Name: "John2", Age: 28}, PartTimeHours: 10.0},
	}

	shiftTimeEmployee := []ShiftEmployee{
		{Employee: Employee{Id: 1, Name: "Akbar", Age: 30}, ShiftHours: 5.0},
		{Employee: Employee{Id: 2, Name: "John2", Age: 28}, ShiftHours: 10.0},
	}

	var employees []EmployeeSalaryCalculator = []EmployeeSalaryCalculator{}

	for _, employee := range fullTimeEmployee {
		employees = append(employees, employee)
	}

	for _, employee := range partTimeEmployee {
		employees = append(employees, employee)
	}

	for _, employee := range shiftTimeEmployee {
		employees = append(employees, employee)
	}

	for _, employee := range employees {
		salary, tax := employee.SalaryCalculate()
		fmt.Printf("\n Employee (%T): %v, Salary: %v, Tax: %v\n", employee, employee, salary, tax)
	}
}

type EmployeeSalaryCalculator interface {
	SalaryCalculate() (salary float64, tax float64)
}

type Employee struct {
	Id   int
	Name string
	Age  int
}

type FullTimeEmployee struct {
	Employee
	ExtraHours float64
}

type PartTimeEmployee struct {
	Employee
	PartTimeHours float64
}

type ShiftEmployee struct {
	Employee
	ShiftHours float64
}

func (e FullTimeEmployee) SalaryCalculate() (Salary float64, tax float64) {
	Salary = e.ExtraHours * ExtraHoursRate * 1.4
	tax = Salary * TaxRate
	Salary -= tax
	return
}

func (e PartTimeEmployee) SalaryCalculate() (Salary float64, tax float64) {
	Salary = e.PartTimeHours * HoursSalary
	tax = Salary * TaxRate
	Salary -= tax
	return 
}

func (e ShiftEmployee) SalaryCalculate() (Salary float64, tax float64) {
	Salary = e.ShiftHours * HoursSalary
	tax = Salary * TaxRate
	Salary -= tax
	return
}
