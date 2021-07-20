package main

import (
	"errors"
	"fmt"
)

// Matrix uses 0-based indexing everywhere except while specifying dimensions
type Matrix struct {
	rows     uint16
	cols     uint16
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

	m := Matrix{uint16(rows), uint16(cols), elements}
	return m, nil
}

func (m *Matrix) GetRows() uint16 {
	return m.rows
}

func (m *Matrix) GetCols() uint16 {
	return m.cols
}

func (m *Matrix) GetElement(row, col int) int {
	return m.elements[row][col]
}

func main() {
	t, _ := NewMatrix(-2, 4)
	fmt.Println(t)
}
