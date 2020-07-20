package main

import "fmt"

func main() {

	// CONTROL FLOW.

	myValue := 21
	// This is a case of a switch without a tag. In this case it doesn't compare
	// something, but it evaluates boolean expressions.
	switch {
	case myValue < 0:
		fmt.Println("My value is negative.")
	case myValue <= 10:
		fmt.Println("My value is between 0 and 10")
	case myValue <= 30: // Will enter here!
		fmt.Println("My value is between 11 and 30")
	default:
		fmt.Println("My variable is greater than 30")
	}

	x := 0
	fmt.Println()
	fmt.Println("First 50 even numbers:")
	// Infinite loop!
	for {
		// You better put something here that can kill the loop.
		if x >= 100 {
			fmt.Printf("%d!", x)
			break // To exit a loop.
		} else {
			fmt.Printf("%d, ", x)
		}
		x += 2
	}
	fmt.Println()

	x = 0 // x was already assigned before, so just overwrite it.
	fmt.Println()
	fmt.Println("First 50 odd numbers:")
	// Infinite loop!
	for {
		// You better put something here that can kill the loop.
		if x%2 == 0 {
			x++
			continue // To skip this loop iteration.
		}
		if x >= 100 {
			fmt.Printf("%d!", x)
			break // To exit a loop.
		} else {
			fmt.Printf("%d, ", x)
		}
		// Increment only one by one, but "even" iterations are
		// skipped using `continue` just for the sake of demonstrating
		// the functionality of that keyword.
		x++
	}
	fmt.Println()

	// USER INPUT SCANNING.

	// Scan reads user input. It takes a pointer as an argument. Data that's
	// inputted by an user is written into the pointer. It returns the number
	// of scanned items ~ tokens that a person typed in (space separated) and
	// also a possible error.
	fmt.Println()
	var appleNum int
	fmt.Printf("Number of apples? ")
	num, err := fmt.Scan(&appleNum)
	// If you don't want the number of tokens, just do _, err := ... to ignore.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Read %d tokens.\n", num)
		fmt.Printf("Got %d apples!\n", appleNum)
	}
}
