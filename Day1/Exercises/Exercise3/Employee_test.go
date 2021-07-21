package main

import "testing"

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
