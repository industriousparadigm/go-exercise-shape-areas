package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{3, 3}
	sq := square{3}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (tr triangle) getArea() float64 {
	return tr.base * tr.height / 2
}

func (sq square) getArea() float64 {
	return math.Pow(sq.sideLength, 2)
}
