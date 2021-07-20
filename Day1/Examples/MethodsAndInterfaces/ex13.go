package main

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func Abs2(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex5{3, 4}
	Scale(v, 10)
	fmt.Println(Abs2(v))
}

