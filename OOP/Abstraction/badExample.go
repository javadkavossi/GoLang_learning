package main

const (
	BaseSalary     = 100
	ExtraHoursRate = 10
	HoursSalary    = 20
	TaxRate        = 0.09
	ShiftSalary    = 30
)

type Employee struct {
	Id            int
	Name          string
	Age           int
	Time          float32
	Type          string
	SalaryOfMonth int
}

func main() {

	employees := []Employee{
		{Id: 1, Name: "PartTime", Age: 30, Time: 10, Type: "Manager", SalaryOfMonth: 50000},
		{Id: 2, Name: "John2", Age: 28, Time: 8, Type: "FullTime", SalaryOfMonth: 50000},
		{Id: 2, Name: "John3", Age: 17, Time: 3, Type: "FullTime", SalaryOfMonth: 18000},
		{Id: 2, Name: "John4", Age: 17, Time: 7, Type: "Shift", SalaryOfMonth: 17000},
	}
	for _, employee := range employees {
		salary, tax, Type := employee.CalculatePartTimeAndFullTimeSalary(employee)
		println(Type, ":", salary, tax)
	}
}

func (e *Employee) FullTimeEmployeeCalculateSalary(extraHours float32) (Salary float32, tax float32) {
	Salary = BaseSalary + (extraHours*ExtraHoursRate)*1.4
	tax = Salary * TaxRate
	Salary -= tax
	return
}

func (e *Employee) PartTimeEmployeeCalculateSalary(hours float32) (Salary float32, tax float32) {
	Salary = hours * HoursSalary
	tax = Salary * TaxRate
	Salary -= tax
	return
}

func (e *Employee) ShiftEmployeeCalculateSalary(hours float32) (Salary float32, tax float32) {
	Salary = hours * HoursSalary
	tax = Salary * TaxRate
	Salary -= tax
	return
}

func (e *Employee) CalculatePartTimeAndFullTimeSalary(Employee Employee) (int, int, string) {
	if Employee.Type == "FullTime" {
		salary, tax := Employee.FullTimeEmployeeCalculateSalary(Employee.Time)
		return int(salary), int(tax), Employee.Type
	} else if Employee.Type == "Shift" {
		salary, tax := Employee.ShiftEmployeeCalculateSalary(Employee.Time)
		return int(salary), int(tax), Employee.Type
	} else {
		salary, tax := Employee.PartTimeEmployeeCalculateSalary(Employee.Time)
		return int(salary), int(tax), Employee.Type
	}
}
