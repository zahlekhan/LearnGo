package main

import "testing"

func TestNewMatrix(t *testing.T) {
	t.Run("Negative or zero values", func(t *testing.T) {
		_,err := NewMatrix(-2,0)
		want := "Indices cannot be negative or zero"
		if err == nil {
			t.Fatalf("expected an error")
		}

		if err.Error() != want {
			t.Errorf("got %q want %q", err, want)
		}
	})
}

func TestGetRows(t *testing.T) {
	t.Run("Create a 1x2 matrix", func(t *testing.T) {
		matrix,err := NewMatrix(1,2)

		if err == nil {
			t.Fatalf("got an error %s",err.Error())
		}

		got := matrix.GetRows()
		want := uint16(1)

		if got != want {
			t.Errorf("got %q want %q",got,want)
		}
	})
}
