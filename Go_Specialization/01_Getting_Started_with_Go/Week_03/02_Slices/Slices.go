package main

import "fmt"

func main() {

	// SLICES

	// A lot of times, slices are used instead of arrays, because they are
	// flexible, you can increase them in size, decrease them in size, etc.
	// A slice is a "window" on an underlying array. There can be an array
	// of lets say 100 values, a slice is just a window of e.g. 3 values.
	// The array could be the same size as the slice, so effectively the
	// slice is just looking at the entire array.

	// Every slice has three properties:
	// 	* Pointer: indicates the start of the slice.
	// 	* Length: is the number of elems in slice.
	// 	* Capacity: the maximum number of elems (from pointer to array's end).

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	// Result: [7]string
	fmt.Printf("%T\n", arr)

	// This is a slicing operation, you'll get a slice/window of positions
	// 1, 2, so the values will be ["b", "c"]
	s1 := arr[1:3]
	// This is a slicing operation, you'll get a slice/window of positions
	// 2, 3, 4, so the values will be ["c", "d", "e"]
	s2 := arr[2:5]
	// Remember, slicing operations ranges are [start, end)
	// Slices can overlap, and refer to the same elements inside the underlying
	// array, so be careful.
	// Result: [b c] [c d e]
	fmt.Println(s1, s2)
	// Changes in the original array.
	arr[2] = "z"
	// Result: [b z] [z d e]
	fmt.Println(s1, s2)
	// Changes in the slices.
	s1[0] = "x"
	s2[0] = "y"
	// Result: [a x y d e f g] [x y] [y d e]
	fmt.Println(arr, s1, s2)

	// This is a slice of the whole array!
	slice1 := arr[:]
	fmt.Println(slice1)

	// This is a slice of only the 1st position.
	slice2 := arr[1:2]
	// Since it's the 1st position, the capacity is the length from there
	// until the end of the underlying array, so only 6 out of 7 elements.
	// Result: []string
	fmt.Printf("%T\n", slice2)
	// Result: Length 1, Capacity 6
	fmt.Printf("Length %d, Capacity %d\n", len(slice2), cap(slice2))

	// When passing to a function, pass slices! Because if you pass an array,
	// it will be passed as VALUE meaning the entirety of it will be copied.
	// Apparently most of Go APIs take slices, not arrays, as arguments.

	// Slices can be initialized directly. When using a slice literal to define
	// a slice, it automatically creates the underlying array, and it covers
	// all of it, slice points to start of array, length of slice is capacity.
	// Notice that there's nothing in the [], compiler says "this is a slice!"
	slice3 := []int{7, 2, 1, 0}
	fmt.Println(slice3)
}
