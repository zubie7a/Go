package main

import "fmt"

// Vertex : a struct to hold coordinates.
type Vertex struct {
	X, Y int
	W, Z int
}

func main() {

	// POINTERS
	i, j := 42, 2701
	// Pointer to `i`.
	p := &i
	// Read `i` through the pointer.
	fmt.Println(*p)
	// Write to `i` through the pointer.
	*p = 21
	// See the new value of `i`.
	fmt.Println(i)

	// Point now to `j`.
	p = &j
	// Edit `j` through the pointer.
	*p = *p / 37
	// See the new value of `j`.
	fmt.Println(j)

	// STRUCTS
	v := Vertex{1, 2, 0, 7}
	v.Z = 4
	fmt.Println(v)
	// And of course, a pointer to a struct.
	pv := &v
	// A struct pointer can access a member directly.
	// (*pv).W = 777
	pv.W = 777
	fmt.Println(v)

	var (
		// Struct initialization with values need all of them,
		// otherwise: `too few values in struct initializer`.
		v1 = Vertex{1, 2, 7, 14}
		// If initializing with literals as 'X' then can pass any number
		// of values or any order, and the rest will default to zero-values.
		v2 = Vertex{X: 1}
		// If initializing with nothing, then all will default to
		// the zero-values.
		v3 = Vertex{}
		// And again, initialization of struct straight into pointer.
		pv4 = &Vertex{22, 222, 222, 2222}
	)
	// {1 2 7 14} {1 0 0 0} {0 0 0 0} &{22 222 222 2222}
	fmt.Println(v1, v2, v3, pv4)

	// Anonymous structs!
	anon := struct {
		i int
		b bool
	}{
		i: 100, b: true,
	}
	fmt.Println("Anonymous struct:", anon)
}
