package main

import "fmt"

func main() {

	// ARRAYS

	// Composite data types put together, aggregate other data types.
	// An array is a fixed-length series of elements of a chosen type.
	// Array elements are accessed using subscript notation, [] like
	// in other languages.

	// When an array is created, elements are initialized to their
	// zero value, which doesn't happen in other languages.
	emptyArray := [3]int{}
	// Result: [0 0 0]
	fmt.Println(emptyArray)

	// This array is initialized with a single value, but 3 cells.
	var oneValArray = [3]int{21}
	// Result: [21 0 0]
	fmt.Println(oneValArray)

	twoValArray := [3]int{14, 42}
	// Result: [14, 42, 0]
	fmt.Println(twoValArray)

	// This line however, would give an error, can't initialize an array
	// with more elements than its capacity.
	// test := [3]int{14, 42, 51, 19}

	var x [5]int
	// Result: [0, 0, 0, 0, 0]
	fmt.Printf("%v\n", x)

	// The size of the array is inferred by the amount of elements provided.
	fourValArray := [...]int{14, 42, 51, 19}
	// Result: [14, 42, 51, 19]
	fmt.Println(fourValArray)

	// Can iterate arrays with this loop and `range` keyword.
	for i, v := range fourValArray {
		// i and v contains the index at value for the array.
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
