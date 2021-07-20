package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	rows uint16
	cols uint16
	elements [][]int
}

func NewMatrix(rows,cols int) (Matrix, error){
	if rows<1 || cols <1 {
		return Matrix{},errors.New("Indices cannot be negative or zero")
	}

	elements := make([][]int,rows)

	for row := range elements {
		elements[row] = make([]int,cols)
	}
	m := Matrix{uint16(rows),uint16(rows),elements}
	return m,nil
}

func main() {
	t,_ := NewMatrix(-2,4)
	fmt.Println(t)
}
