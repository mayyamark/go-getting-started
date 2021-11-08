package main

import "fmt"

func main() {
	introDemo()
	fmt.Println("-----------------")

	skipingDemo()
}

func introDemo() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		/*
			2**0 = 1
			2**1 = 2
			2**2 = 4
			2**3 = 8
			2**4 = 16
			2**5 = 32
			2**6 = 64
			2**7 = 128
		*/
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func skipingDemo() {
	pow := make([]int, 10)
	// Use the index only
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// skip the index, use the value only
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
