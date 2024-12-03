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
		{Id: 1, Name: "Ali", Age: 30, ExtraHours: 10.0},
		{Id: 2, Name: "John2", Age: 28, ExtraHours: 5.0},
	}

	partTimeEmployee := []PartTimeEmployee{
		{Id: 1, Name: "Akbar", Age: 30, PartTimeHours: 5.0},
		{Id: 2, Name: "John2", Age: 28, PartTimeHours: 10.0},
	}

	shiftTimeEmployee := []ShiftEmployee{
		{Id: 1, Name: "Akbar", Age: 30, ShiftHours: 5.0},
		{Id: 2, Name: "John2", Age: 28, ShiftHours: 10.0},
	}

	var employees []Employee = []Employee{}

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

	// var employees Employee = fullTimeEmployee[0]
	// salary, tax := employees.SalaryCalculate()
	// fmt.Println(employees, salary, tax)

	// employees = fullTimeEmployee[1]
	// salary, tax = employees.SalaryCalculate()
	// fmt.Println(employees, salary, tax)

	// employees = partTimeEmployee[0]
	// salary, tax = employees.SalaryCalculate()
	// fmt.Println(employees, salary, tax)

	// employees = partTimeEmployee[1]
	// salary, tax = employees.SalaryCalculate()
	// fmt.Println(employees, salary, tax)

	// employees = shiftTimeEmployee[0]
	// salary, tax = employees.SalaryCalculate()
	// fmt.Println(employees, salary, tax)

	// employees = shiftTimeEmployee[1]
	// salary, tax = employees.SalaryCalculate()
	// fmt.Println(employees, salary, tax)

}

type Employee interface {
	SalaryCalculate() (salary float64, tax float64)
}

type FullTimeEmployee struct {
	Id         int
	Name       string
	Age        int
	ExtraHours float64
}

type PartTimeEmployee struct {
	Id            int
	Name          string
	Age           int
	PartTimeHours float64
}

type ShiftEmployee struct {
	Id         int
	Name       string
	Age        int
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
