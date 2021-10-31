package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// %T => prints the type	%v => prints the value
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// The expression T(v) converts the value v to the type T.
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // or f := float64(...)
	var u uint = uint(f)                          // or u := uint(...)
	fmt.Println(x, y, u)
}
