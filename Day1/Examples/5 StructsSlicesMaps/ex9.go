package main

import "fmt"

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex2{1, 2}  // has type Vertex
	v2 = Vertex2{X: 1}  // Y:0 is implicit
	v3 = Vertex2{}      // X:0 and Y:0
	p  = &Vertex2{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

