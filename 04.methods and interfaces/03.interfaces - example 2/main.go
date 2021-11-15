package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.
type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
// Here we implement geometry on rects.
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * r.height * r.width
}

// The implementation for circles.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, then we can call methods that are in the named interface.
// Here’s a generic measure function taking advantage of this to work on any geometry
func measure(g geometry) {
	fmt.Println("geometry:", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimeter", g.perimeter())
	fmt.Println("--------------")
}

func main() {
	r := rectangle{height: 2, width: 5}
	c := circle{radius: 3}

	// The circle and rect struct types both implement the geometry interface
	// so we can use instances of these structs as arguments to measure.
	measure(r)
	measure(c)
}
