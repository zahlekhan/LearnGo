package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func Abs(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex4{3, 4}
	fmt.Println(Abs(v))
}

