package employee

// inte ha medlemsfunktioner i sig
type Employee struct {
	Name          string // mandatory
	Age           int    // mandatory
	StreetAddress string
	PostalCode    int
	City          string
}

func New(name string, age int) Employee {
	return Employee{Name: name, Age: age}
}

func (emp Employee) CalculateSalary() int { // Detta blir en medlemsmetod för Employee
	return emp.Age * 100
}

// func CalculateSalary(employee Employee) int {
// 	return employee.Age * 100
// }
