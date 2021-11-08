package main

import (
	"fmt"
	"strings"
)

func main() {
	introDemo()
	fmt.Println("-----------------")

	mutatingDemo()
	fmt.Println("-----------------")

	lengthAndCapacityDemo()
	fmt.Println("-----------------")

	zeroValueDemo()
	fmt.Println("-----------------")

	buildInMakeDemo()
	fmt.Println("-----------------")

	tikTacToeDemo()
	fmt.Println("-----------------")

	appendDemo()
}

func introDemo() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // select from index 1 to index 4
	fmt.Println(s)            // [3, 5, 7]
}

func mutatingDemo() {
	characters := [4]string{
		"A",
		"B",
		"C",
		"D",
	}
	fmt.Println(characters) // [A B C D]

	selectedCharacters := characters[1:3]

	selectedCharacters[0] = "XXX"               // we are mutating the characters array as well!
	fmt.Println(selectedCharacters, characters) // [XXX C] [A XXX C D]
}

func lengthAndCapacityDemo() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("s", s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice("s", s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice("s", s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice("s", s) // len=2 cap=4 [5 7]

}

func zeroValueDemo() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func buildInMakeDemo() {
	a := make([]int, 5)
	printSlice("a", a) // a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // d len=5 cap=3 [0 0 0]
}

func tikTacToeDemo() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendDemo() {
	var s2 []int
	printSlice("s2", s2) // s2 len=0 cap=0 []

	// append works on nil slices.
	s2 = append(s2, 0)
	printSlice("s2", s2) // s2 len=1 cap=1 [0]

	// The slice grows as needed.
	s2 = append(s2, 1)
	printSlice("s2", s2) // s2 len=2 cap=2 [0 1]

	// We can add more than one element at a time.
	s2 = append(s2, 2, 3, 4)
	printSlice("s2", s2) // s2 len=5 cap=6 [0 1 2 3 4 5]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
