package main

import (
	"fmt"
)

// Copy your Sqrt function from the earlier exercise and modify it to return an error value.

// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	stringValue := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %v", stringValue)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)

	for i := 1; i <= 10000000000; i++ {
		zCached := z

		if z -= (z*z - x) / (2 * z); zCached/1000 <= z/1000 {
			fmt.Println(z, zCached)
			break
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
