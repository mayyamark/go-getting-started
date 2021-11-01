package main

import (
	"fmt"
	"math"
)

// Implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.
func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 1; i <= 10000000000; i++ {
		zCached := z

		// The formula				 if it's not the first operation (since we guess that we start with z = 1)
		if z -= (z*z - x) / (2 * z); i != 1 && int(zCached*1000) <= int(z*1000) {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
