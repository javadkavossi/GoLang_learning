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

	var employees Employee = fullTimeEmployee[0]
	salary, tax := employees.SalaryCalculate(float64(fullTimeEmployee[0].ExtraHours))
	fmt.Println(employees, salary, tax)

	employees = fullTimeEmployee[1]
	salary, tax = employees.SalaryCalculate(float64(fullTimeEmployee[1].ExtraHours))
	fmt.Println(employees, salary, tax)

	employees = partTimeEmployee[0]
	salary, tax = employees.SalaryCalculate(float64(partTimeEmployee[0].PartTimeHours))
	fmt.Println(employees, salary, tax)

	employees = partTimeEmployee[1]
	salary, tax = employees.SalaryCalculate(float64(partTimeEmployee[1].PartTimeHours))
	fmt.Println(employees, salary, tax)

	employees = shiftTimeEmployee[0]
	salary, tax = employees.SalaryCalculate(float64(shiftTimeEmployee[0].ShiftHours))
	fmt.Println(employees, salary, tax)

	employees = shiftTimeEmployee[1]
	salary, tax = employees.SalaryCalculate(float64(shiftTimeEmployee[1].ShiftHours))
	fmt.Println(employees, salary, tax)

}

type Employee interface {
	SalaryCalculate(hours float64) (salary float64, tax float64)
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

func (e FullTimeEmployee) SalaryCalculate(extraHours float64) (Salary float64, tax float64) {
	Salary = extraHours * ExtraHoursRate * 1.4
	tax = Salary * TaxRate
	Salary -= tax
	return
}

func (e PartTimeEmployee) SalaryCalculate(hours float64) (Salary float64, tax float64) {
	Salary = hours * HoursSalary
	tax = Salary * TaxRate
	Salary -= tax
	return
}

func (e ShiftEmployee) SalaryCalculate(hours float64) (Salary float64, tax float64) {
	Salary = hours * HoursSalary
	tax = Salary * TaxRate
	Salary -= tax
	return
}
