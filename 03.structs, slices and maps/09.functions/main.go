package main

import (
	"fmt"
	"math"
)

func main() {
	introDemo()
	fmt.Println("-----------------")

	closuresDemo()
	fmt.Println("-----------------")
}

func introDemo() {
	// Functions are values too. They can be passed around just like other values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// Function values may be used as function arguments and return values.
	fmt.Println(hypot(5, 12))      // 5*5 + 12*12 = 169 => sqrt(169) = 13
	fmt.Println(compute(hypot))    // 3*3 + 4*4 = 25 => sqrt(25) = 5
	fmt.Println(compute(math.Pow)) // 3^4 = 81
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func closuresDemo() {
	// Go functions may be closures (closure - a function value that references variables from outside its body).
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}
}

// The adder function returns a closure. Each closure is bound to its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
