package main

import (
	"encoding/json"
	"errors"
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

func (m *Matrix) SetElement(row, col, value int) error {
	if row < 0 || row >= m.GetRows() || col < 0 || col >= m.GetCols() {
		return errors.New("indices are out of bounds")
	}
	m.elements[row][col] = value
	return nil
}

func (m *Matrix) ToJSONString() (string, error) {
	res, err := json.Marshal(m.elements)

	if err != nil {
		return "", err
	}

	return string(res), nil
}

func AddMatrix(m1, m2 *Matrix) (Matrix, error) {
	if m1.GetRows() != m2.GetRows() && m1.GetCols() != m2.GetCols() {
		return Matrix{}, errors.New("dimensions don't match")
	}

	res, _ := NewMatrix(m1.GetRows(), m1.GetCols())

	for row := 0; row < m1.GetRows(); row++ {
		for col := 0; col < m1.GetCols(); col++ {
			element1, _ := m1.GetElement(row, col)
			element2, _ := m2.GetElement(row, col)
			_ = res.SetElement(row, col, element1+element2)
		}
	}

	return res, nil
}
