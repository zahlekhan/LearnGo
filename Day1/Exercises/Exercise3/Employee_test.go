package main

import (
	"testing"
)

func TestContractor_IncrementTotalDuration(t *testing.T) {
	emp, _ := EmployeeFactory("John", "Full time", 500)
	emp.IncrementTotalDuration(30)
	got := emp.GetTotalDurationWorked()
	want := 30

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestEmployee_Salary(t *testing.T) {
	emp, _ := EmployeeFactory("John", "Full time", 500)
	emp.IncrementTotalDuration(30)
	got := emp.Salary()
	want := 15000

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestTotalWagesOwed(t *testing.T) {
	emp1, _ := EmployeeFactory("John", "Full time", 500)
	emp1.IncrementTotalDuration(30)

	emp2, _ := EmployeeFactory("Alice", "Freelancer", 10)
	emp2.IncrementTotalDuration(300)

	got := TotalWagesOwed([]Employee{emp1, emp2}...)
	want := 18000

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
