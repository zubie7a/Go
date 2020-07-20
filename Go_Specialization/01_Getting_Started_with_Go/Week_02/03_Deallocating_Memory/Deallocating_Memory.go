package main

import "fmt"

// These declarations go into the heap. This a persistent memory region.
// Data on the heap must be deallocated when it is done being used, in most
// compiled languages this is done manually.
// In C, to allocate in the heap you'd need to do `x = malloc(32);` and then
// to release that memory you'd need to do `free(x);`.
var z = 42

func f() {
	// When a variable is no longer needed, it should be deallocated.
	// Otherwise, we'll eventually run out of memory. Each call to f()
	// creates an integer.

	// Declarations inside functions go into the stack. The stack is an
	// area of memory primarily dedicated to function calls. In regular
	// languages, stack memory is deallocated once the function completes.
	// That's why tail recursion is good, because then the compiler can
	// optimize and not keep every frame of the stack unnecessarily.

	// In Go this is different however...
	var x = 4
	fmt.Printf("%d\n", x)
}

func main() {
}
