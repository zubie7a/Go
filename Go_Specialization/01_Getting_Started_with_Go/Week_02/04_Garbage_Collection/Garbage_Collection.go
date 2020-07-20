package main

import "fmt"

// These declarations go into the heap. This a persistent memory region.
// Data on the heap must be deallocated when it is done being used, in most
// compiled languages this is done manually.
// In C, to allocate in the heap you'd need to do `x = malloc(32);` and then
// to release that memory you'd need to do `free(x);`.
var z = 42

func f() *int {
	// When a variable is no longer needed, it should be deallocated.
	// Otherwise, we'll eventually run out of memory. Each call to f()
	// creates an integer.

	// Declarations inside functions go into the stack. The stack is an
	// area of memory primarily dedicated to function calls. In regular
	// languages, stack memory is deallocated once the function completes.
	// That's why tail recursion is good, because then the compiler can
	// optimize and not keep every frame of the stack unnecessarily.

	// In Go this is different however... You can declare a variable in a
	// function, but return a reference to it! Because of this reference,
	// the variable won't be deallocated.
	var x = 4

	return &x
}

func main() {

	// GARBAGE COLLECTION.

	var y *int
	// Got the pointer from `x` inside f() and stored it in `y`. Because now
	// that pointer is being used here, `x` won't be deallocated. Go tracks
	// existing references to variables before deallocating them.
	y = f()
	fmt.Println(*y)
	// Once this function finishes, the pointer to `x` won't be used anymore,
	// so Garbage Collector will deallocate it. ONLY when all refs are gone.
}
