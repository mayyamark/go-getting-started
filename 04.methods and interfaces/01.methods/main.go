package main

import (
	"fmt"
	"math"
)

func main() {
	introDemo()
	fmt.Println("---------------")

	myFloatTypeDemo()
	fmt.Println("---------------")

	pointerReceiverDemo()
	fmt.Println("---------------")
	fmt.Println("---------------")
}

// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.

// The receiver appears in its own argument list between the func keyword and the method name.

// In this example, the Abs method has a receiver of type Vertex named v.
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func introDemo() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// You can declare a method on non-struct types, too.

// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int).

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func myFloatTypeDemo() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

// For example, the Scale method here is defined on *Vertex.

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

// With a value receiver, the Scale method would operate on a copy of the original Vertex value.
// (This is the same behavior as for any other function argument.)
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc - rewritern Scale, ordinary function, it takes a pointer as an argument
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerReceiverDemo() {
	v := Vertex{3, 4}
	v.Scale(10) // Methods with pointer receivers take either a value or a pointer as the receiver when they are called => v.Scale(10) as (&v).Scale(10)

	// ScaleFunc(&v, 10) ... while functions with a pointer argument must take a pointer

	fmt.Println(v.Abs())

	p := &v
	fmt.Println(p.Abs()) // Methods with value receivers take either a value or a pointer as the receiver when they are called => p.Abs() is interpreted as (*p).Abs()
}

// 2 reaasonst to use a pointer receiever:
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
