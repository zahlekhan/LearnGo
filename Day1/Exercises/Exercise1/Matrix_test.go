package main

import "testing"

func TestNewMatrix(t *testing.T) {
	t.Run("Negative or zero values", func(t *testing.T) {
		_, err := NewMatrix(-2, 0)
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
		matrix, err := NewMatrix(1, 2)

		if err != nil {
			t.Fatalf("got an error %s", err.Error())
		}

		got := matrix.GetRows()
		want := 1

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestGetCols(t *testing.T) {
	t.Run("Create a 1x2 matrix", func(t *testing.T) {
		matrix, err := NewMatrix(1, 2)

		if err != nil {
			t.Fatalf("got an error %s", err.Error())
		}

		got := matrix.GetCols()
		want := 2

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestGetElement(t *testing.T) {
	t.Run("Get element 0,1 from 1x2 matrix", func(t *testing.T) {
		matrix, initErr := NewMatrix(1, 2)

		if initErr != nil {
			t.Fatalf("got an error %s", initErr.Error())
		}

		got, _ := matrix.GetElement(0, 1)
		want := 0

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Get element 1,2 from 1x2 matrix", func(t *testing.T) {
		matrix, initErr := NewMatrix(1, 2)

		if initErr != nil {
			t.Fatalf("got an error %s", initErr.Error())
		}

		_, err := matrix.GetElement(1, 2)
		want := "indices are out of bounds"

		if err == nil {
			t.Fatalf("expected an error")
		}

		if err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	})
}

func TestSetElement(t *testing.T) {
	t.Run("Set element 0,1 to 3 in a 1x2 matrix", func(t *testing.T) {
		matrix, initErr := NewMatrix(1, 2)

		if initErr != nil {
			t.Fatalf("got an error %s", initErr.Error())
		}

		setErr := matrix.SetElement(0, 1, 3)

		if setErr != nil {
			t.Fatalf("got an error %s", initErr.Error())
		}

		got, _ := matrix.GetElement(0, 1)
		want := 3

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Set element 0,1 to 3 in a 1x2 matrix", func(t *testing.T) {
		matrix, initErr := NewMatrix(1, 2)

		if initErr != nil {
			t.Fatalf("got an error %s", initErr.Error())
		}

		err := matrix.SetElement(1, 2, 3)
		want := "indices are out of bounds"

		if err == nil {
			t.Fatalf("expected an error")
		}

		if err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	})
}
