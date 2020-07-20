package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	// STRING PACKAGES.

	// Unicode:
	// Mostly useful for manipulating individual runes. Characters retrieved
	// from the string are of type byte, so you need to cast them to rune()
	// before calling the functions of unicode package.
	myString := "xyz721!X🔥"

	// %c with myString[3] will print a "7".
	// %b with myString[3] will print a "110111", which is 55 in binary,
	// because in ASCII/UTF-8, the character at position 55 is "7".
	// Use %t to print a boolean.
	fmt.Printf("%c is a digit? %t\n",
		myString[3], unicode.IsDigit(rune(myString[3])))

	// %c with myString[2] will print a "z".
	// %b with myString[2] will print a "1111010", which is 122 in binary,
	// because in ASCII/UTF-8, the character at position 122 is "z".
	// Use %t to print a boolean.
	fmt.Printf("%c is a digit? %t\n",
		myString[2], unicode.IsDigit(rune(myString[2])))

	fmt.Printf("%c is a letter? %t\n",
		myString[2], unicode.IsLetter(rune(myString[2])))

	fmt.Printf("%c is a lower case? %t\n",
		myString[7], unicode.IsLower(rune(myString[7])))

	fmt.Printf("%c is a upper case? %t\n",
		myString[7], unicode.IsUpper(rune(myString[7])))

	fmt.Printf("%c is punctuation? %t\n",
		myString[6], unicode.IsPunct(rune(myString[6])))

	// STRINGS PACKAGE:
	// Useful for operating on the whole string, not individual runes.

	// Compare(a, b) : returns an integer comparing two strings lexicographically.
	// 	* a == b: result is 0.
	// 	* a < b: result is -1.
	// 	* a > b: result is  1.
	firstName, middleName := "Santiago", "Martin"
	shortName := firstName[:5]

	// S... is lexicographically higher than M...
	fmt.Println("S... is higher than M...",
		strings.Compare(firstName, middleName))

	// "Santiago" contains "go" :-)
	fmt.Printf("%s contains go? %t\n",
		firstName, strings.Contains(firstName, "go"))

	// Check prefixes.
	fmt.Printf("%s is the prefix of %s? %t\n",
		shortName, firstName, strings.HasPrefix(firstName, shortName))

	// Replace two occurences of a string in a given string.
	fmt.Printf("Replace 2 'a' by 'x'? %s\n",
		strings.Replace(firstName+middleName, "a", "x", 2))

	// Replace all occurences of a string in a given string.
	fmt.Printf("Replace all 'a' by 'x'? %s\n",
		strings.ReplaceAll(firstName+middleName, "a", "x"))

	// Similarly you can make ToLower on the whole string.
	fmt.Printf("WHY ARE YOU YELLING, %s?\n", strings.ToUpper(firstName))

	fileLine := "This line was read from a file\n"
	// Uh, this will print a new line because Println, but also another newline
	// because the line read from file was not sanitized.
	fmt.Println(fileLine)
	// This will trim unnecessary whitespaces from the line read from file!
	fmt.Println(strings.TrimSpace(fileLine))

	strInt := "72142"
	badStrInt := "a72142"
	// Try converting a string to an integer.
	if intStr, err := strconv.Atoi(strInt); err == nil {
		fmt.Printf("%.2f\n", float32(intStr))
	}
	// There will be an error if the string is not an integer.
	if _, err := strconv.Atoi(badStrInt); err == nil {
		fmt.Println("Conversion successed!")
	} else {
		// This will print the error.
		fmt.Printf("%v\n", err)
		// Can do it either way...
		// fmt.Println(err)
	}

	intVal := 72142
	// newStr := "a" + intVal // This will give an error.
	// When converting int to string there's no possible format problems.
	newStr := "z" + strconv.Itoa(intVal)
	fmt.Println(newStr)

	// FormatFloat (f, fmt, prec, bitSize):
	// 	* converts floating point number to a string.
	// ParseFloat (s, bitSize):
	// 	* Converts a string to a floating point number.
}
