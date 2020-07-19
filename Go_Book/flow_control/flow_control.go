package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func getOS() string {
	// Short initializers are also supported in `switch` in its scope.
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	default:
		// freebsd, openbsd, plan9, windows:
		return os
	}
}

func getGreeting() string {
	t := time.Now()
	// Can have a `switch` without a variable which implies 'true' and then can do anything
	// in any of the cases, easier to do this than a long list of if ... else if ... else if...
	switch {
	case t.Hour() < 12:
		return "Good Morning!"
	case t.Hour() < 17:
		return "Good Afternoon!"
	default:
		return "Good Evening!"
	}
}

func pow(x, n, lim float64) float64 {
	// `if` blocks can have a short statement to declare variable(s) only for the scope.
	// of the `if`, just like `for`.
	v := math.Pow(x, n)
	if v < lim {
		return v
	}

	fmt.Printf("%g >= %g\n", v, lim)
	return lim
}

// Sqrt : Computes sqrt of a float
func Sqrt(x float64) float64 {
	z := 1.0
	limit := 10
	fmt.Printf("Approximating the square root of %v with %v iterations\n", int(x), limit)
	for i := 0; i < limit; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func theFinalCountdown() {
	// Defers literally 'defers' the execution of something until the surrounding function returns.
	// You can also stack them together as in a LIFO!
	// This whole function is pretty much written backwards, so this will print...
	// Countdown!
	// 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
	// BOOM!
	defer fmt.Println("BOOM!")
	for i := 0; i < 10; i++ {
		if i > 0 {
			defer fmt.Printf(", ")
		} else {
			defer fmt.Printf("\n")
		}
		defer fmt.Printf("%v", i)
	}
	defer fmt.Println("Countdown!")
}

func main() {
	fmt.Println("Hello Universe,")
	fmt.Println(getGreeting())
	fmt.Printf("Go is running on %s.\n", getOS())

	fmt.Println("Nice and shiny squares all around")
	// Traditional `for` syntax.s
	for i := 0; i < 10; i++ {
		fmt.Printf("%v", i*i)
		if i+1 < 10 {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}

	k := 0
	sum := 0
	// There's no `while` but can do a `for` only with condition.
	// or `for ; k <= 100 ; { ... }
	for k <= 100 {
		sum += k
		k++
	}
	fmt.Println("Sum from 1 to 100 is... !")
	fmt.Println(sum)
	// INFINITE LOOP DO NOT USE UNLESS IN GOROUTINE.
	// for {
	// 	fmt.Println("YOLO")
	// }

	fmt.Println(pow(3531515143141, 2, 10))
	fmt.Println(pow(3431434234234, 3, 20))
	Sqrt(729)
	theFinalCountdown()
}
