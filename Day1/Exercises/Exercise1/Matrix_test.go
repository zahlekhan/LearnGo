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
