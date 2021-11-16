package main

import "fmt"

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and assign value to v.
// (The data flows in the direction of the arrow.)

// Like maps and slices, channels must be created before use: ch := make(chan int)

// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// The example code sums the numbers in a slice, distributing the work between two goroutines.
// Once both goroutines have completed their computation, it calculates the final result.

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // [1, 2, 3] => 1 + 2 + 3 = 6
	go sum(s[len(s)/2:], c) // [4, 5, 6] => 4 + 5 + 6 = 15
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y) // 5	15	21
}
