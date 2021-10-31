package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// The expression need not be surrounded by parentheses ( ) but the braces { } are required.
	if v := math.Pow(x, n); v < lim { // The if statement can start with a short statement to execute before the condition.
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // Variables declared inside an if short statement are also available inside any of the else blocks.
	}
	// Can't use v here!
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
