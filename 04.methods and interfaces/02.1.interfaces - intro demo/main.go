package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
type I interface {
	M()
}

// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.
// Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
