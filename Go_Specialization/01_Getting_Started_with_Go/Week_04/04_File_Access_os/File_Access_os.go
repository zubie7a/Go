package main

import (
	"fmt"
	"os"
)

func fileRead(f *os.File, bArr []byte) {
	// `f.Read` will fill the byte array until it's full or until the
	// source file is completely read.

	numBytes, err := f.Read(bArr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Read %d bytes. \"%s\"\n", numBytes, string(bArr))
	}
}

func fileWrite(f *os.File, bArr []byte) {

	numBytes, err := f.Write(bArr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Wrote %d bytes. \"%s\"\n", numBytes, string(bArr))
	}
}

func fileWriteStr(f *os.File, str string) {

	numBytes, err := f.WriteString(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Wrote %d bytes. \"%s\"\n", numBytes, str)
	}
}

func main() {

	// FILE ACCESS `os`

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

	// `os` has some more advanced file access functions to have more
	// control over file read/write operations.

	fRead, err := os.Open("test.txt")
	// Result: *os.File
	fmt.Printf("%T\n", fRead)
	// Call `close` of `fRead` at the end of this function.
	defer fRead.Close()
	if err != nil {
		fmt.Println(err)
	} else {

		// Number of bytes we want to read.
		nb := 10
		// Create a byte array of `nb` bytes.
		bArr := make([]byte, nb)
		// Result: []byte
		fmt.Printf("%T\n\n", bArr)

		// Result: Read 10 bytes. "Hello, my "
		fileRead(fRead, bArr)
		// Result: Read 10 bytes. "name is Sa"
		fileRead(fRead, bArr)
		// Result: Read 10 bytes. "ntiago!
		//
		// I"
		fileRead(fRead, bArr)

		// If you notice, the original byte array size was quite low, only
		// 10 bytes. Every read will read 10 bytes only, and resume from
		// the place the last read stopped (where the read head is).
		// If the original byte array was quite huge, let's say 500 bytes,
		// then the whole test file would fit into it and read all 100 bytes.
	}

	fmt.Println()

	fWrite, err := os.Create("outfile.txt")
	// Result: *os.File
	fmt.Printf("%T\n", fWrite)
	// Call `close` of `fWrite` at the end of this function.
	defer fWrite.Close()

	bArr2 := []byte{76, 79, 76, 10}
	fileWrite(fWrite, bArr2)
}
