package main

import "fmt"

// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

func main() {
	// A struct literal denotes a newly allocated struct value by listing the values of its fields.
	v1 := Vertex{1, 2}
	var v2 = Vertex{X: 1}  // We can specify the fields	=>	Y:0 is implicit
	var v3 = Vertex{}      // X:0 and Y:0
	var p1 = &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, v2, v3, p1)

	// Struct fields are accessed using a dot.
	v1.X = 4
	fmt.Println(v1.X)

	p2 := &v1 // Struct fields can be accessed through a struct pointer.
	p2.Y = 3  // We could write (*p).Y, but Go permits us to do it without the explicit dereference.
	fmt.Println(v1, p2, *p2)

}
