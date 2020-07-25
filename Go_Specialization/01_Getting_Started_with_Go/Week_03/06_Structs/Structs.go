package main

import "fmt"

func main() {

	// STRUCTS

	// Structs are another aggregate data type, groups together other objects
	// of arbitrary type.

	type Person struct {
		name string
		city string
		mail string
		cash int
	}

	p1 := Person{
		name: "Santiago",
		city: "Amsterdam",
		mail: "S@Z.com",
		// cash: 77777, // If not provided, will initialize to 0.
	}

	// Result: {Santiago Amsterdam S@Z.com}
	fmt.Println(p1)

	p2 := Person{}
	// This struct was not provided values, so it will initialize with the
	// default values (empty string and 0).
	// Result: {   0}
	fmt.Println(p2)

	// A struct can also be initialized without the keys, but in this case
	// need to provide all the fields, otherwise it won't compile.
	p3 := Person{"Martin", "Medellin", "M@Z.com", 77777}
	// Result: {Martin Medellin M@Z.com 77777}
	fmt.Println(p3)

	// This seems to be effectively the same than p4 := Person{}
	// So the default value is a struct with all values default inside?
	var p4 Person
	// Result: {   0}
	fmt.Println(p4)

	// Result: main.Person
	fmt.Printf("%T\n", p4)
	// Struct fields have to be accessed with dot notation.
	p4.name = "SM"
	p4.city = "AMS"
	p4.mail = "SM@ZO.com"
	p4.cash = 2142
	// Result: {SM AMS SM@ZO.com 2142}
	fmt.Printf("%v\n", p4)

	// Struct can also be created with `new`, fields will be initialized
	// to zero-value. This will create a struct BUT return a pointer to it.
	p5 := new(Person)
	p5.cash = 80
	// Result: *main.Person
	fmt.Printf("%T\n", p5)

	// Yet another way to create a pointer from an anonymous struct.
	p6 := &Person{
		cash: 20000,
	}
	p6.cash += 2

	// When accessing the fields, it doesn't matter if it's a struct-pointer or
	// a struct.
	totalCash := p1.cash + p2.cash + p3.cash + p4.cash + p5.cash + p6.cash
	// Result: 100001
	fmt.Println(totalCash)
}
