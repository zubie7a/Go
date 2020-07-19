package main

import (
	"fmt"
	"math"
)

// Global variable that can hold a function.
var mult func(float64, float64) float64

// Function that receives a function, other parameters, and evaluates.
func compute(f func(float64, float64) float64, a, b int) float64 {
	return f(float64(a), float64(b))
}

// Function that returns a function.
func getHypFunc() func(float64, float64) float64 {
	hyp := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}
	return hyp
}

// A closure to a function that every time that its called returns the
// next fibonacci value.

// The variables are only in the scope of the function(s), which can
// access and assign to them, only these functions, no other can access.
// If fibClosure() is called twice, the returned functions will be on separate
// closures, but the two functions returned by a single call of this function
// are on the same closure.
func fibClosure() (func() int, func() string) {
	a, b := 1, 0
	return func() int {
			// First function returned, it generates the next fibonacci value.
			a, b = b, b+a
			return a
		}, func() string {
			// Second function returned, it tells what the previous and
			// current values are, like inspecting state of closure.
			return fmt.Sprintf("Cur: %v, Next %v", a, b)
		}
}

func main() {
	// Assigning something to the global function.
	mult = func(a, b float64) float64 {
		return a * b
	}
	// Functions are values. They can be passed around, they can be returned,
	// they can even be assigned to variables.
	hyp := getHypFunc()

	fmt.Println(hyp(3, 4))
	fmt.Println(compute(hyp, 1, 2))
	fmt.Println(compute(math.Pow, 1, 2))
	fmt.Println(mult(6, 7))

	fib, inspector := fibClosure()
	for i := 0; i < 21; i++ {
		fmt.Println(fib())
		fmt.Println(inspector())
	}

	// Next: https://tour.golang.org/methods/1
}
