package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	Name         string
	EmployeeType string
	Wages
}

type Wages interface {
	GetDurationUnit() string
	GetTotalDurationWorked() int
	GetRate() int
	IncrementTotalDuration(int)
}

// Salary can be refactored to be included in Wages interface if wage policies get too complex
func (e *Employee) Salary() int {
	return e.GetRate() * e.GetTotalDurationWorked()
}

func EmployeeFactory(name, employeeType string, rate int) (Employee, error) {
	var wagesFactory Wages
	switch employeeType {
	case "Full time":
		wagesFactory = &fullTime{rate, 0}
	case "Contractor":
		wagesFactory = &contractor{rate, 0}
	case "Freelancer":
		wagesFactory = &freelancer{rate, 0}
	default:
		return Employee{}, errors.New("undefined employee type")
	}
	return Employee{name, employeeType, wagesFactory}, nil
}

type fullTime struct {
	rate                int
	totalDurationWorked int
}

func (f *fullTime) GetDurationUnit() string {
	return "Days"
}

func (f *fullTime) GetTotalDurationWorked() int {
	return f.totalDurationWorked
}

func (f *fullTime) GetRate() int {
	return f.rate
}

func (f *fullTime) IncrementTotalDuration(increment int) {
	f.totalDurationWorked += increment
}

type contractor struct {
	rate                int
	totalDurationWorked int
}

func (c *contractor) GetDurationUnit() string {
	return "Days"
}

func (c *contractor) GetTotalDurationWorked() int {
	return c.totalDurationWorked
}

func (c *contractor) GetRate() int {
	return c.rate
}

func (c *contractor) IncrementTotalDuration(increment int) {
	c.totalDurationWorked += increment
}

type freelancer struct {
	rate                int
	totalDurationWorked int
}

func (f *freelancer) GetDurationUnit() string {
	return "Hours"
}

func (f *freelancer) GetTotalDurationWorked() int {
	return f.totalDurationWorked
}

func (f *freelancer) GetRate() int {
	return f.rate
}

func (f *freelancer) IncrementTotalDuration(increment int) {
	f.totalDurationWorked += increment
}

func main() {
	emp1, _ := EmployeeFactory("John", "Full time", 500)
	emp1.IncrementTotalDuration(30)
	fmt.Println(emp1.GetTotalDurationWorked())
	fmt.Println(emp1.Salary())
}
