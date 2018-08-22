package main

import (
	"fmt"
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
	newTriangle := triangle{
		height: 15,
		base:   4,
	}

	newSquare := square{
		sideLength: 20,
	}

	printArea(newTriangle)
	printArea(newSquare)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
