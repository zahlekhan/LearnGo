package main

import (
	"errors"
	"fmt"
)

// Matrix uses 0-based indexing everywhere except while specifying dimensions
type Matrix struct {
	rows     int
	cols     int
	elements [][]int
}

func NewMatrix(rows, cols int) (Matrix, error) {
	if rows < 1 || cols < 1 {
		return Matrix{}, errors.New("Indices cannot be negative or zero")
	}

	// Initialize slice with given size
	elements := make([][]int, rows)
	for row := range elements {
		elements[row] = make([]int, cols)
	}

	m := Matrix{rows, cols, elements}
	return m, nil
}

func (m *Matrix) GetRows() int {
	return m.rows
}

func (m *Matrix) GetCols() int {
	return m.cols
}

func (m *Matrix) GetElement(row, col int) (int, error) {
	if row < 0 || row >= m.GetRows() || col < 0 || col >= m.GetCols() {
		return 0, errors.New("indices are out of bounds")
	}
	return m.elements[row][col], nil
}

func main() {
	t, _ := NewMatrix(-2, 4)
	fmt.Println(t)
}
