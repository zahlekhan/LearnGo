package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat1(-math.Sqrt2)
	v := Vertex9{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat1 float64

func (f MyFloat1) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex9 struct {
	X, Y float64
}

func (v *Vertex9) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
