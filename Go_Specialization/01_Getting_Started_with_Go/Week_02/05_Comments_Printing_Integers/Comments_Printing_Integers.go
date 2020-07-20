package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	// COMMENTS.
	// This is a comment.
	var x int // This is another comment.
	/*
		This is a block comment.
	*/
	x++ // Default value of int is 0, increment will make it 1.
	fmt.Println(x)

	// PRINTING.
	name := "Santi"
	fmt.Printf("Hi %s!\n", name)
	fmt.Println("How are you " + name + "?")

	now := time.Now()
	fmt.Println("Current time is " + now.String())
	fmt.Println("UNIX format is " + now.Format(time.UnixDate))
	// Go uses some very specific constants to format dates...
	// https://golang.org/src/time/format.go
	// https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html
	fmt.Println("YYYY-MM-DD is " + now.Format("2006-01-01"))
	fmt.Println("DD-MM-YYYY is " + now.Format("01-01-2006"))
	fmt.Println("YYYY-MM-DD is " + now.Format("2006-01-01"))
	fmt.Println("YYYY-MM-DD hh:mm:ss is " + now.Format("2006-01-02 15:04:05"))
	// https: //gobyexample.com/epoch
	fmt.Printf("Epoch seconds: %d\n", now.Unix())
	fmt.Printf("Epoch nanoseconds %d\n", now.UnixNano())

	numFloat := 27 / 22.0 // Type inference for float needs one to be float, or will be an int.
	// Print a float, specifying the number of decimal places, similar to C/C++.
	fmt.Printf("%.14f\n", numFloat)

	// INTEGERS.
	var (
		// int8: from -128 to 127
		min8, max8 int8
		// int16: from -32768 to 32767
		min16, max16 int16
		// int32: from -2147483648 to 2147483647
		min32, max32 int32
		// int64: from -9223372036854775808 to 9223372036854775807
		min64, max64 int64
		// uint8: from 0 to 255
		umin8, umax8 uint8
		// uint16: from 0 to 65535
		umin16, umax16 uint16
		// uint32: from 0 to 4294967295
		umin32, umax32 uint32
		// uint64: from 0 to 18446744073709551615
		umin64, umax64 uint64
	)

	// Signed integers.
	min8, max8 = math.MinInt8, math.MaxInt8
	min16, max16 = math.MinInt16, math.MaxInt16
	min32, max32 = math.MinInt32, math.MaxInt32
	min64, max64 = math.MinInt64, math.MaxInt64
	// Unsigned integers.
	umin8, umax8 = 0, math.MaxUint8
	umin16, umax16 = 0, math.MaxUint16
	umin32, umax32 = 0, math.MaxUint32
	umin64, umax64 = 0, math.MaxUint64

	fmt.Printf("int8: from %d to %d\n", min8, max8)
	fmt.Printf("int16: from %d to %d\n", min16, max16)
	fmt.Printf("int32: from %d to %d\n", min32, max32)
	fmt.Printf("int64: from %d to %d\n", min64, max64)
	fmt.Printf("uint8: from %d to %d\n", umin8, umax8)
	fmt.Printf("uint16: from %d to %d\n", umin16, umax16)
	fmt.Printf("uint32: from %d to %d\n", umin32, umax32)
	fmt.Printf("uint64: from %d to %d\n", umin64, umax64)

	// Integers also have these binary operators:
	// 	* Arithmetic: + - * / % << >>
	// 	* Comparisson: == != > < >= <=
	//  * Boolean: && ||
}
