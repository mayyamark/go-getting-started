package main

import "fmt"

func main() {
	introDemo()
	fmt.Println("-----------------")

	mapLiteralsDemo()
	fmt.Println("-----------------")

	workingWithMapsDemo()
}

type Vertex struct {
	Lat, Long float64
}

func introDemo() {
	// A map maps keys to values.
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	// The make function returns a map of the given type, initialized and ready for use.
	var m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967}]
}

func mapLiteralsDemo() {
	// Map literals are like struct literals, but the keys are required.
	var m = map[string]Vertex{
		"Bell Labs": { // or 		"Bell Labs": Vertex{...}
			40.68433, -74.39967,
		},
		"Google": { // or 		"Google": Vertex{...}
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

}

func workingWithMapsDemo() {
	m := make(map[string]int)

	// Insert or update an element in map m: m[key] = elem
	m["Answer"] = 42

	// Retrieve an element: elem = m[key]
	fmt.Println("The value:", m["Answer"]) // The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) // The value: 48

	// Delete an element: delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // The value: 0 (zero value)

	// Test that a key is present with a two-value assignment: elem, ok := m[key]
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}
