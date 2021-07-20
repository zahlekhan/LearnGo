package main

import "fmt"

type I3 interface {
	M()
}

func main() {
	var i I3
	describe3(i)
	i.M()
}

func describe3(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}

