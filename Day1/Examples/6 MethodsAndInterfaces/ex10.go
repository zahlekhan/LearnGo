package main

import (
	"fmt"
"math"
)

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex3{3, 4}
	fmt.Println(v.Abs())
}


