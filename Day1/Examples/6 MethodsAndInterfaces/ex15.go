package main

import (
	"fmt"
	"math"
)

type Vertex7 struct {
	X, Y float64
}

func (v Vertex7) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex7) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex7{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex7{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

