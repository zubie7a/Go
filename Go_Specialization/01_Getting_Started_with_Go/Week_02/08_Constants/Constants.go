package main

import "fmt"

// Constant: a expressiom whose value is known at compile time.
// Type is inferred from righthand side (boolean, string, number)
const (
	// Inferred `float64`
	x = 1.3
	// Inferred `int`
	y = 4
	// Inferred `string`
	z = "Hi"
	// Explicit `float32`
	w float32 = 2.4
)

func main() {

	// CONSTANTS.

	// iota: a function to generate constants. It generates a set of
	// related but distinct constants. Often represents a property
	// which has several distinct possible values. This is also known
	// as a one-hot. You know a variable it's going to be one-hot coded
	// if it can have one to five values, each one is a distinct constant.

	// One example is days of the week, each day can be represented by
	// a distinct constant, but they are related. A variable to represent
	// days of the week then can have only 7 possible distinct values.
	// The actual value doesn't really matter, what matters is that the
	// values are different so we are able to distinguish the constant,
	// so every constant with iota must have a distinct value.
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)

	// By using iota, each constant is assigned to an unique integer, but
	// the actual value inside of it doesn't really matter. The implementation
	// starts at 1 and increments one by one.
	type DayOfWeek int
	const (
		MON DayOfWeek = iota
		TUE
		WED
		THU
		FRI
		SAT
		SUN
	)

	// Will print `0 1 2 3 4 5`.
	fmt.Println(A, B, C, D, E, F)
	// Will print `0 1 2 3 4 5 6`.
	fmt.Println(MON, TUE, WED, THU, FRI, SAT, SUN)
	// Will print `float64 int string float32`
	fmt.Printf("%T %T %T %T\n", x, y, z, w)
	// Will print `main.Grades main.DayOfWeek`
	fmt.Printf("%T %T\n", A, MON)
}
