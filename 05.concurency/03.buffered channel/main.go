package main

import "fmt"

func main() {
	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
	ch := make(chan int, 2)
	// Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.
	ch <- 1
	// fmt.Println(<-ch) // fatal error: all goroutines are asleep - deadlock!
	ch <- 2
	// ch <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
