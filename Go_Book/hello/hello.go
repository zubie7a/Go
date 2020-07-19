package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

/*
Unlike in C, in Go assignment between items of different type requires an explicit conversion.
Go's basic types are:

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
    	 // represents a Unicode code point

	float32 float64

	complex64 complex128

	according to precission, (int, float64, complex128) are auto inferred.
		* i := 1     (int)
		* f := 1.0   (float64)
	    * c := 2i    (complex128)

The zero values are:

	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.

*/

// Constants are typeless and if numeric, take type according to its context.
// Can have super high precision, the problem is converting them into a type
// of lower precision.
const (
	// ./hello.go:111:41: constant 1267650600228229401496703205376 overflows uint64
	Big   = 1 << 100
	Small = Big >> 99
	Strc  = "STR_CONST"
)

func int10x(x int) int {
	return x * 10
}

func float10x(x float64) float64 {
	return x * 10.0
}

func str10x(x string) string {
	return x + " 10!"
}

// By default false.
// Outside of a function can't use the := short assignment.
var c, python, java bool

var (
	// ToBe : a boolean variable.
	ToBe = false
	// MaxInt : the max value that can be held in uint64.
	MaxInt uint64 = 1<<64 - 1
	// A complex value.
	cz = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	// Can return tuples and stuff.
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// Named return, don't have to declare named variables and will be returned in the end.
	return
}

func main() {
	fmt.Println("Hello Universe,")

	fmt.Println("My favorite number is", add(7, add(14, 21)), "±", rand.Intn(21))
	fmt.Println("The square root of 729 is", math.Sqrt(729))
	fmt.Println("The value of Pi is ", math.Pi)
	// Initialize multiple values.
	s, m, z := "Santiago", "Martin", 10
	_s, _m := swap(s, m)
	fmt.Println(_s, _m, z)
	fmt.Println(split(17))
	// By default 0.
	var i int
	// By default "".
	var str string
	// Can be initialized this way
	// x, y := 1, 2 // or...
	var x, y int = 1, 2
	fmt.Println(i, str, c, python, java, x, y)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", float10x(Small), float10x(Small))
	fmt.Printf("Type: %T Value: %v\n", int10x(Small), int10x(Small))
	fmt.Printf("Type: %T Value: %v\n", str10x(Strc), str10x(Strc))
}

/*
	General:

		%v	the value in a default format
			when printing structs, the plus flag (%+v) adds field names
		%#v	a Go-syntax representation of the value
		%T	a Go-syntax representation of the type of the value
		%%	a literal percent sign; consumes no value

	Boolean:

		%t	the word true or false

	Integer:

		%b	base 2
		%c	the character represented by the corresponding Unicode code point
		%d	base 10
		%o	base 8
		%q	a single-quoted character literal safely escaped with Go syntax.
		%x	base 16, with lower-case letters for a-f
		%X	base 16, with upper-case letters for A-F
		%U	Unicode format: U+1234; same as "U+%04X"

	Floating-point and complex constituents:

		%b	decimalless scientific notation with exponent a power of two,
			in the manner of strconv.FormatFloat with the 'b' format,
			e.g. -123456p-78
		%e	scientific notation, e.g. -1.234456e+78
		%E	scientific notation, e.g. -1.234456E+78
		%f	decimal point but no exponent, e.g. 123.456
		%F	synonym for %f
		%g	%e for large exponents, %f otherwise. Precision is discussed below.
		%G	%E for large exponents, %F otherwise

	String and slice of bytes (treated equivalently with these verbs):

		%s	the uninterpreted bytes of the string or slice
		%q	a double-quoted string safely escaped with Go syntax
		%x	base 16, lower-case, two characters per byte
		%X	base 16, upper-case, two characters per byte

	Pointer:

		%p	base 16 notation, with leading 0x

	The default format for %v is:

		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p

	For compound objects, the elements are printed using these rules, recursively, laid out like this:

		struct:             {field0 field1 ...}
		array, slice:       [elem0 elem1 ...]
		maps:               map[key1:value1 key2:value2]
		pointer to above:   &{}, &[], &map[]
*/
