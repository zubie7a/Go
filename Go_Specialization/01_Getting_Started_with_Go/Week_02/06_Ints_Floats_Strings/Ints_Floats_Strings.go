package main

import "fmt"

func main() {

	// INTEGERS.

	// Most binary operators need operands of the same type...
	// Go does not do implicit type conversion?
	var x int32 = 1
	var y int16 = 2
	// x = y // This line would give an error.
	x = int32(y) // Instead, have to cast it!
	x++
	fmt.Println(x) // Will print 3.

	// FLOATING POINTS.

	// float32 : ~6 digits of precision.
	// float64 : ~15 digits of precision.
	var f1 float32
	var f2 float64
	f1 = 27 / 22.0
	f2 = 27 / 22.0
	fmt.Println(f1) // Will print 1.2272727
	fmt.Println(f2) // Will print 1.2272727272727273

	// COMPLEX VALUES.

	// There's `complex64` and `complex128`, the later is inferred if using
	// the complex() function, if you want specificly a `complex64` you need
	// to specify the type.
	var z complex64 = complex(2, 3)
	fmt.Println(z) // Will print (2+3i)

	var w complex64 = complex(3, 4)
	fmt.Println(z * w) // Will print (-6+17i)
	fmt.Println(z + w) // Will print (5+7i)
	fmt.Println(z / w) // Will print (0.72+0.04i)
	// Complex number multiplication:
	// 	(x + yi)(u + vi) = (x*u - y*v) + (x*v + y*u)i
	// 	(2 + 3i)(3 + 4i) = (2*3 - 3*4) + (2*4 + 3*3)i = (-6+17i)

	// STRINGS.

	// They are ASCII and Unicode. They are sequences of bytes.

	// Character coding: each character is associated with an (7) 8-bit number:
	// 	* 'A' = 0x41 ...
	// ASCII is 8-bit value, so it can maximum represent 256 different chars,
	// or more like 128 because one of the bits is used for something else...
	// 8-bit is sufficient for English and numbers and common symbols, but
	// once you need to include extended characters like Chinese or Japanese
	// you can't use 8-bits.

	// Unicode is a 32-bit value. UTF-8 is variable length, so 8-bit UTF codes
	// are the same as ASCII codes.
	unicode := "五😍"
	fmt.Println(unicode) // This will actually print the string above :-)
	// 😁 := "😁" Thankfully Go doesn't support emoji variables.
	五 := "😁" // However it does support variables in Japanese
	fmt.Println(五)
	// In UTF-8/Unicode, a "code point" refers to a unicode character.
	// In Go, a "Rune" refers to a "code point". Strings are sequences of
	// arbitrary bytes. However, they are READ ONLY! You can't modify a string,
	// however, you can make a modified version of an existing string into a
	// new string.
	greeting := "Hi how are you"
	name := "Santi"
	// Slices: on a sequence, [x:y] will grab characters in the range [x,y).
	// If x is not given, will do it from start, if y is not given, will do it
	// until the end. [:3] takes characters at positions 0,1,2. [3:] takes the
	// characters starting at position 3 until the end.
	// We have created a new string, without modifying the original strings.
	newGreeting := greeting[:3] + name + ", " + greeting[3:] + "?"
	fmt.Println(newGreeting) // Will print "Hi Santi, how are you?"
}
