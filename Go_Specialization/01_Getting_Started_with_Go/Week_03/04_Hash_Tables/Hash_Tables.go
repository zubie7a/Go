package main

import "fmt"

func main() {

	// HASH TABLES

	// Allows fast access to large amounts of data.
	// Contains Key-Value pairs. Each value is associated with a unique key.
	// There is a hash function and it is used to take a key and compute a
	// "slot" in which to put the value.

	// Advantages:
	// Faster lookup than lists, constant-time vs linear-time.
	// Arbitrary keys, not ints like slices or arrays.

	// Disadvantages:
	// May have collisions, two keys hash to the same slot. There's some
	// ways to handle collisions though, but all is done internally.
	// Even so, collisions may make things a bit slower.

	// In Go, Maps are an implementation of a hash table.
	// Create a map that maps from string to int.
	idMap := make(map[string]int)
	// Assign values with these keys.
	idMap["Santiago"] = 21
	idMap["Zubieta"] = 700
	// Result: 21 700 721
	fmt.Printf("%d %d %d\n",
		idMap["Santiago"], idMap["Zubieta"], idMap["Santiago"]+idMap["Zubieta"])
}
