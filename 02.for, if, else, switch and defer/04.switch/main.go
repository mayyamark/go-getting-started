package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintCurrentOS() {
	fmt.Print("Go runs on ")

	//  Go only runs the selected case, not all the cases that follow
	switch os := runtime.GOOS; os {
	case "darwin":
		// The break statement is provided automatically in Go
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func PrintWhenIsSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday { // Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	case today + 4:
		fmt.Println("In four days.")
	default:
		fmt.Println("Too far away.")
	}
}

func PrintGreetinf() {
	t := time.Now()

	switch { // Switch without condition. This construct can be a clean way to write long if-then-else chains.
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	PrintCurrentOS()
	fmt.Println("-----------------")
	PrintWhenIsSaturday()
	fmt.Println("-----------------")
	PrintGreetinf()
}
