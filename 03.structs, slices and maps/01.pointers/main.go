package main

import "fmt"

// A pointer holds the memory address of a value.
func main() {
	i, j := 42, 168

	// The & operator generates a pointer to its operand.
	p := &i        // point to i
	fmt.Println(p) // the address of i (e.g., 0xc0000b8008)

	// The * operator denotes the pointer's underlying value. This is known as "dereferencing" or "indirecting":
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 2    // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
