package main

import "fmt"

func main() {

	// POINTERS.

	var x = 721 // Can omit type since it will be inferred from right-hand side.
	var y int
	var ip *int // `ip` is a pointer to int type.

	ip = &x        // `&` takes the address pointer to x, then assign it to ip.
	y = *ip        // '*' dereferences the pointer, and assigns the value to y.
	fmt.Println(y) // the result should be the value in x, 721.

	// `new` creates a variable and returns a pointer to that variable.
	// this way doing short assignment to pointers is possible.
	ptr := new(int)
	// Assign 3 to the variable by dereferencing the pointer.
	*ptr = 3
}
