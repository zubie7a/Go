package main

import "fmt"

// This is a "global" variable, and it's in the scope of all functions.
// Remember that at this level you can't do short assignment.
var x = 7

func f() {
	// This declaration of `x` is inside a block and completely contained,
	// so any later call to `x` will resolve to the global scope.
	{
		x := 42
		x++
	}
	// This will use the global declaration of `x`.
	fmt.Printf("%d\n", x)
}

func g(x int) {
	// The global declaration of `x` is ignored because this function has an
	// `x` parameter.
	fmt.Printf("%d\n", x+1)
}

func h() {
	// The global declaration of `x` is ignored because this function has a
	// redeclaration of `x`.
	x := 1
	fmt.Printf("%d\n", x)
}

func main() {
	f()  // Will print 7.
	g(1) // Will print 2.
	h()  // Will print 1.
}

// Go is lexically scoped using blocks.
// Implicit blocks:
//     * Universe block: all Go source.
//     * Package block: all source in a package.
//     * File block: all source in a single file.
//     * "if", "for", "switch": all code inside statement.
//     * Clauses in "switch" or "select": individual clauses get a block.
