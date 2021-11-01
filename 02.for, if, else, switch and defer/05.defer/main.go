package main

import "fmt"

func PrintHelloWorld() {
	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer fmt.Println("World")

	fmt.Println("Hello")
}

func PrintCounter() {
	fmt.Println("Counting...")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i) // Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	}

	fmt.Println("Done!")
}

func main() {
	PrintHelloWorld()
	fmt.Println("-----------------")
	PrintCounter()
}
