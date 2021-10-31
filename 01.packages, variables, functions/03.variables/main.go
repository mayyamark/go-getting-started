package main

import "fmt"

var a, b int = 1, 2 // The var statement declares a list of variables; as in function argument lists, the type is last. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

func main() {
	var c int // A var statement can be at package or function level.
	fmt.Println(a, b, c)

	var d, e int = 1, 2
	f := 3 // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type. When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	e = 4
	fmt.Println(d, e, f)

	// Variables declared without an explicit initial value are given their zero value.
	var i int
	var b bool
	var s string
	fmt.Printf("%v %v %v\n", i, b, s)
}
