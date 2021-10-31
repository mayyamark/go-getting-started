package main

import ( // This code groups the imports into a parenthesized, "factored" import statement.

	"fmt"
	"math"
	"math/rand" // By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi) // In Go, a name is exported if it begins with a capital letter. For example Pi is exported from the math package.
}
