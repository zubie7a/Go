package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// FILE ACCESS `ioutil`

	// It is true in all languages that file access is linear, not random.
	// But a file might be in a random access device, like RAM or a SSD,
	// but it is accessed as if it was linear.

	// Basic Operations
	// 	1. Open - get handle for access.
	// 	2. Read - read bytes into []byte.
	// 	3. Write - write []byte into file.
	// 	4. Close - release handle.
	// 	5. Seek - move read/write head.

	// There's more than one package in Go that allows to access files, it
	// has functions in it that allows file access.

	// `ioutil` has some basic file access functions.

	// This path has to be at the same level the command is run.
	// Explicit open/close are not needed with this method.
	data1, err := ioutil.ReadFile("test.txt")
	if err != nil {
		// If running at the wrong level...
		// Result: open test.txt: no such file or directory
		fmt.Println(err)
	} else {
		// The contents of the whole file goes into this []byte variable.
		// Because of this, large files will cause a problem, you'll take
		// a lot of memory if not all. Use `ioutil.ReadFile` for small files,
		// and also if you don't have too many of them, loading a lot of small
		// files completely into memory can still make it run out.
		fmt.Println(string(data1))
	}

	// Convert a string to a byte array.
	data2 := []byte("This is a test writing a file.\nGood luck!\n")
	// Write this byte array into a file. UNIX-style permission bytes.
	err = ioutil.WriteFile("outfile.txt", data2, 0777)
}
