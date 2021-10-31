package main

import "fmt"

/*
	The basic for loop has three components separated by semicolons:
		- the init statement: executed before the first iteration
		- the condition expression: evaluated before every iteration
		- the post statement: executed at the end of every iteration
	There are no parentheses surrounding the three components of the for statement and the braces { } are always required.
*/
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 { // The init and post statements are optional <=> while
		sum2 += sum2
	}
	fmt.Println(sum2)
}
