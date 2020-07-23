package main

import "fmt"

func main() {

	// VARIABLE SLICES

	// There's a function called `make` and can be used to make slices. It
	// allows making other things, but for now focus on slices.

	// We first talked previously about creating first an underlying array,
	// and then creating a slice out of it. Then we talked about creating a
	// slice directly with a slice literal with certain values, which created
	// the underlying array automatically.

	// There's a third way though. The slice is created directly, but not
	// initializing with any particular values. You want to initialize a
	// slice to a particular size, without giving it any values, at least
	// at the beginning.

	// Two argument version:
	// Specify type, and length/capacity.
	slice1 := make([]int, 10)
	// Result:
	// 	Type: []int
	// 	Values: [0 0 0 0 0 0 0 0 0 0]
	// 	Length: 10
	// 	Capacity 10
	fmt.Printf("Type: %T\nValues: %v\nLength: %v\nCapacity %v\n",
		slice1, slice1, len(slice1), cap(slice1))

	fmt.Println()

	// Three argument version:
	// Specify type, length, and capacity separately.
	// Size is 5, but capacity is 10, so we can increase up to size 10.
	slice2 := make([]int, 5, 10)
	// Result:
	// 	Type: []int
	// 	Values: [0 0 0 0 0]
	// 	Length: 5
	// 	Capacity 10
	fmt.Printf("Type: %T\nValues: %v\nLength: %v\nCapacity %v\n",
		slice2, slice2, len(slice2), cap(slice2))

	fmt.Println()

	// Size of a slice can be increased by `append()`
	// This adds elements to the end of a slice. This is done by
	// putting the element in the underlying array, and increases
	// the capacity of the slice by n-added-elements.

	arr := [7]int{7, 14, 21, 42, 127, 721, 777}
	// Result: [7 14 21 42 127 721 777]
	fmt.Println(arr)
	slice3 := arr[2:5]
	// Result:
	// 	Type: []int
	// 	Values: [21 42 127]
	// 	Length: 3
	// 	Capacity 5
	fmt.Printf("Type: %T\nValues: %v\nLength: %v\nCapacity %v\n",
		slice3, slice3, len(slice3), cap(slice3))

	fmt.Println()

	// Now let's try "appending" something to the slice.
	slice3 = append(slice3, 420)
	// Result:
	// 	Type: []int
	// 	Values: [21 42 127 420]
	// 	Length: 4
	// 	Capacity 5
	fmt.Printf("Type: %T\nValues: %v\nLength: %v\nCapacity %v\n",
		slice3, slice3, len(slice3), cap(slice3))

	// Notice that the slice now has length 4! But still capacity 5.
	// Check the status of the underlying array though...
	// Result: [7 14 21 42 127 420 777]
	fmt.Println(arr)
	// Appending to the slice overwrote the value at the same
	// position in the array!!!

	fmt.Println()

	// What happens if we do it once more?
	slice3 = append(slice3, 999)
	// Result:
	// 	Type: []int
	// 	Values: [21 42 127 420 999]
	// 	Length: 5
	// 	Capacity 5
	fmt.Printf("Type: %T\nValues: %v\nLength: %v\nCapacity %v\n",
		slice3, slice3, len(slice3), cap(slice3))
	// Result: [7 14 21 42 127 420 999]
	fmt.Println(arr)

	fmt.Println()

	// Now we've reached the moment of truth. The slice length reached
	// it's capacity, what happens if we try appending to it once more?
	slice3 = append(slice3, 1000)
	// Result:
	// 	Type: []int
	// 	Values: [21 42 127 420 999 1000]
	// 	Length: 6
	// 	Capacity 10
	fmt.Printf("Type: %T\nValues: %v\nLength: %v\nCapacity %v\n",
		slice3, slice3, len(slice3), cap(slice3))
	// The capacity of the slice is double what it previously was!
	// The previous underlying array remains the same... This is because
	// once the slice reached capacity, it stopped looking at this array,
	// and a new underlying array was created as a copy of the original
	// one, with twice the capacity!
	fmt.Println(arr)

	fmt.Println()

	// By creating a new slice starting at the slice until capacity,
	// we can see the new underlying array.
	// Result: [21 42 127 420 999 1000 0 0 0 0]
	fmt.Println(slice3[:cap(slice3)])

	// So, this is why it's good to create a slice with a fairly good
	// initial capacity, so extending it is really inserting into the
	// underlying array, and only when the capacity of that array is
	// reached then a new underlying array will be allocated.
}
