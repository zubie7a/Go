package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// CONTROL FLOW.

	x := 0
	// Simple condition.
	if x <= 10 {
		fmt.Printf("The value of x is %d!\n", x)
	}

	strInt := "72142"
	badStrInt := "a72142"
	// Try converting a string to an integer.
	// A condition that first declares variables only useable inside the
	// if block, and then evaluates a condition on one of those variables.
	if intStr, err := strconv.Atoi(strInt); err == nil {
		fmt.Printf("%.2f\n", float32(intStr))
	}

	// There will be an error if the string is not an integer.
	if _, err := strconv.Atoi(badStrInt); err == nil {
		fmt.Println("Conversion successed!")
	} else {
		// A regular else like in any other language.
		// This will print the error.
		fmt.Printf("%v\n", err)
		// Can do it either way...
		// fmt.Println(err)
	}

	fmt.Println()
	fmt.Println("First 50 even numbers:")
	// Infinite loop!
	for {
		// You better put something here that can kill the loop.
		if x >= 100 {
			fmt.Printf("%d!", x)
			break
		} else {
			fmt.Printf("%d, ", x)
		}
		x += 2
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("First 20 powers of two:")
	// Regular loop.
	// for <init>; <cond>; <update>
	for x := 0; x < 20; x++ {
		fmt.Printf("%d ", int64(math.Pow(2.0, float64(x))))
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("First 20 fibonacci numbers:")
	// Declare a slice to store fibonacci numbers, initialized
	// with two values. Some reading on slices vs arrays:
	// https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
	fib := []int{0, 1}
	// Start at position 2.
	idx := 2
	// A loop structured like a `while`, Go does not have `while`.
	for idx < 20 {
		// At every iteration look two steps back.
		// f(n) = f(n-2) + f(n-1)
		fibVal := fib[idx-2] + fib[idx-1]
		// Append the new value to the slice. Unlike arrays, slices can be
		// resized using the built-in `append` function.
		fib = append(fib, fibVal)
		// Increase the position index.
		idx++
	}
	fmt.Println(fib)

	fmt.Println()
	// The `switch` is a multi-way if statement. You don't need to do "break",
	// if you want it to continue you have to explicitly mention `fallthrough`
	// It may have a <tag> which is a variable to compare in each case.
	myValue := 21
	switch myValue {
	case 20:
		fmt.Println("The variable is 20")
	case 21: // Will enter here!
		fmt.Println("The variable is 21")
		// fallthrough
	default:
		fmt.Println("Unknown value")
	}
}
