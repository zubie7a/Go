package main

import (
	"fmt"
	"strings"
)

func main() {
	// ARRAYS
	// Must be declared with a fixed length, as their nature is static.
	var a [2]string
	a[0] = "Hello"
	a[1] = "Universe"
	// ["Hello" "Universe"]
	fmt.Printf("%q\n", a)

	fib := [9]int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	fmt.Println(fib)

	// SLICES
	// Dynamic length storage, sort of references to arrays.
	var newFib []int
	// This is a static array.
	base := [2]int{0, 1}
	// A slice with half-open range [l,r).
	//   * [l:r]   : from position l to r-1, if r == l + 1 then will take only one.
	//   * [:r]    : from position 0 to r-1.
	//   * [l:]    : from position l to len-1.
	//   * [:]     : the whole array from 0 to len-1.
	//   Sadly, can't use "-n" at the end to implicitly do len-n.
	numSli := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("From 2 to 6:\t%v\n", numSli[2:6])
	fmt.Printf("From _ to 4:\t%v\n", numSli[:4])
	fmt.Printf("From 5 to _:\t%v\n", numSli[5:])

	baseSlice := base[0:2]
	// Slices can be appended, and unpacked (...) separates all values of slice.
	newFib = append(newFib, baseSlice...)
	for i := 1; i < 21; i++ {
		f_ni := newFib[i] + newFib[i-1]
		// Now sure how efficient is to keep appending to end of list...
		newFib = append(newFib, f_ni)
	}

	fmt.Printf("Fibonacci_21:\t%v\n", newFib)
	// Since slices are "references" to arrays, they do not store anything, any change
	// is done on the underlying array, also affecting other slices pointing to it.
	otherSlice := base[:]
	anotherSlice := base[:]
	yetAnotherSlice := base[:]
	// Oh... we appended something to this slice. Something tells me the underlying array
	// won't be modified, and instead a whole new array will be created?
	yetAnotherSlice = append(yetAnotherSlice, -1)
	// The underlying array was modified.
	fmt.Printf("Original array:\t%v\n", base)
	otherSlice[0] = 777
	fmt.Printf("Modified slice:\t%v\n", otherSlice)
	// The underlying array was modified.
	fmt.Printf("Affected array:\t%v\n", base)
	// Other slices pointing to it reflect the change.
	fmt.Printf("Affected slice:\t%v\n", anotherSlice)
	// Untouched, because it made a resizing operation it created a new array beneath.
	// "If the backing array is too small to fit all the given values a bigger array
	// will be allocated. The returned slice will point to the newly allocated array."
	fmt.Printf("Not affected:\t%v\n", yetAnotherSlice)
	fmt.Printf("Not affected:\t%v\n", newFib)

	// Now, what happens if we get an slice that is smaller than the underlying array,
	// and we append something to it? The slice will grow, and the appended element will really
	// overwrite the element thats next in the underlying array.
	newBase := []int{1, 2, 3}
	fmt.Printf("Another array:\t%v\n", newBase)
	smallSlice := newBase[0:1]
	fmt.Printf("Smaller slice:\t%v\n", smallSlice)
	smallSlice = append(smallSlice, 22)
	fmt.Printf("Modified slice:\t%v\n", smallSlice)
	fmt.Printf("Modified array:\t%v\n", newBase)

	// A slice literal is like an array but without the length.
	arr1 := [3]bool{true, true, false}
	// This will allocate an array just like the one above and then refer to it.
	// So to create slices don't have to first create an array and then slice it,
	// can create it immediately like this.
	sli1 := []bool{true, true, false}
	// This will show that both are "the same".
	fmt.Println()
	fmt.Println(arr1, sli1)

	// Anonymous struct slices!
	anonSlice := []struct {
		i, j int
		b    bool
	}{
		{1, 2, false},
		{3, 4, true},
		{b: true},
		{j: 777},
		{i: -1, b: true},
	}
	fmt.Println("Anonymous slice:", anonSlice, "\n")

	// In slices, must not confuse length and capacity.
	// 	* len(s): the number of elements this slice contains.
	//  * cap(s): the number of elements the underlying array contains,
	//            from the start of the slice.
	// This means that if the capacity of a slice is bigger than its length,
	// appending to it will cause the underlying array element at the next
	// position to be overwritten. If the capacity is the same, then appending
	// will reallocate a new array. Also cap can never be less than len (???)
	arr := [6]int{2, 3, 5, 7, 11, 13}
	s1 := arr[:]
	sx := (s1[2:3])[:0]
	s2 := arr[:3]
	s3 := arr[3:]
	fmt.Printf("Original array:\t%v\n", arr)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)
	printSlice(sx)
	// Even if this slice had zero length, the underlying capacity still allows to
	// redefine its bounds.
	sx = sx[:3]
	printSlice(sx)
	sx = sx[:cap(sx)]
	printSlice(sx)
	// But if you try to slice it beyond capacity then you get in trouble.
	// panic: runtime error: slice bounds out of range
	// sx = sx[:cap(sx) + 1]

	// Default value of a slice is `nil`. Also len and cap are 0, meaning there's
	// no underlying array. Declaring with var w/o assigning makes it have a zero value.
	var nilSli []int
	// Creating this way doesn't give a "zero value".
	nilSli2 := []int{}
	// Both are technically len 0, cap 0, [], etc...
	if nilSli == nil {
		fmt.Printf("nil!\t%v\n", nilSli)
	}
	// ... but only one is considered nil, the other was created explicity to be like this.
	if nilSli2 == nil {
		fmt.Printf("nil!\t%v\n", nilSli2)
	}

	// DYNAMICALLY SIZED ARRAYS!
	// The odd way would be to create an empty slice and start appending stuff to it until
	// reaching capacity. All elements will have zero values. Just do this:
	lenSli := 5
	dyn := make([]int, lenSli)
	printSlice(dyn)
	// That created a slice of len 5, which in turn allocates an array of len 5.
	// Can also specify capacity, so that an array of given capacity is created, but the
	// slice is created only up to the given length.
	capSli := 10
	// Having a big underneath capacity allows to do appends while overwriting the next
	// position without necessarily reaching the point of reallocation. Obviously too big
	// of a capacity is bad for memory, but a just good enough size will allow you do to
	// appends without reallocating a new array every time.
	dyn2 := make([]int, lenSli, capSli)
	printSlice(dyn2)

	fmt.Println("\nTIC TAC TOE MAD TRIQUI TIME!")
	// Multi array drifting! Or how to make slices of other types (including slices)
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns!
	// board[Y][X]
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	// len(board) will give the length of the first dimension.
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Printf("\n%v\n", numSli)
	numSli = []int{7, 6, 5, 4, 3, 2, 1}
	// range is like a foreach of slices and maps. It returns the index and a copy of the value.
	//  _, v := ... to just get the value.
	//     i := ... to just get the index.
	for i, v := range numSli {
		fmt.Printf("At pos %d there's %d\n", i, v)
	}

	// Powers of 2.
	limit := 10
	fmt.Printf("\nThe first %d powers of 2 are:\n", limit)
	pow := make([]int, limit)
	for i := range pow {
		// logical shift '<<' requires unsigned int.
		pow[i] = 1 << uint(i) // == 2**i
	}
	printSlice(pow)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d\t%v\n", len(s), cap(s), s)
}
