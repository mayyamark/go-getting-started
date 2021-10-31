package main

import "fmt"

func add(x int, y int) int { // For the params - the type comes after the variable name.
	return x + y
}

func swap(x, y string) (string, string) { // A function can return any number of results.
	// When two or more consecutive named function parameters share a type, we can omit the type from all but the last
	return y, x
}

func split(sum int) (x, y int) { // Return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values.
	x = sum * 4 / 9
	y = sum - x
	return // A return statement without arguments returns the named return values. This is known as a "naked" return.
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
}
